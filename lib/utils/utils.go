package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Panic is the same as the built-in panic.
var Panic = func(v interface{}) { panic(v) }

// E if the last arg is error, panic it.
func E(args ...interface{}) []interface{} {
	err, ok := args[len(args)-1].(error)
	if ok {
		Panic(err)
	}
	return args
}

// Mkdir makes dir recursively.
func Mkdir(path string) error {
	return os.MkdirAll(path, 0o775)
}

// MustToJSONBytes encode data to json bytes.
func MustToJSONBytes(data interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	E(enc.Encode(data))
	b := buf.Bytes()
	return b[:len(b)-1]
}

// OutputFile auto creates file if not exists, it will try to detect the data type and
// auto output binary, string or json.
func OutputFile(p string, data interface{}) error {
	dir := filepath.Dir(p)
	_ = Mkdir(dir)

	var bin []byte

	switch t := data.(type) {
	case []byte:
		bin = t
	case string:
		bin = []byte(t)
	case io.Reader:
		f, _ := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o664)
		_, err := io.Copy(f, t)
		return err
	default:
		bin = MustToJSONBytes(data)
	}

	return ioutil.WriteFile(p, bin, 0o664)
}
