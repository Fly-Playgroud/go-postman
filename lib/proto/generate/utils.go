package main

import (
	"fmt"
	"github.com/Fly-Playgroud/go-postman/lib/utils"
	"github.com/ysmood/gson"
	"io/ioutil"
	"net/http"
	"strings"
)

func mapType(n string) string {
	return map[string]string{
		"boolean": "bool",
		"number":  "float64",
		"integer": "int",
		"string":  "string",
		"binary":  "[]byte",
		"object":  "map[string]gson.JSON",
		"any":     "gson.JSON",
	}[n]
}

func enumList(schema gson.JSON) []string {
	var enum []string
	if schema.Has("enum") {
		enum = []string{}
		for _, v := range schema.Get("enum").Arr() {
			if _, ok := v.Val().(string); !ok {
				panic("enum type error")
			}
			enum = append(enum, v.Str())
		}
	}

	return enum
}

func parseSchemaVersion(schema schema) string {
	members := strings.Split(schema.String(), "/")
	return members[5]
}

func getSchema(schema schema) gson.JSON {
	res, err := http.Get(schema.String())
	utils.E(err)
	defer func() { _ = res.Body.Close() }()

	data, err := ioutil.ReadAll(res.Body)
	utils.E(err)

	obj := gson.New(data)
	utils.E(utils.OutputFile(fmt.Sprintf("tmp/%s_schema.json", schema.Version()), obj.JSON("", "  ")))
	return obj
}

func enumList(schema gson.JSON) []string {
	var enum []string
	if schema.Has("enum") {
		enum = []string{}
		for _, v := range schema.Get("enum").Arr() {
			if _, ok := v.Val().(string); !ok {
				panic("enum type error")
			}
			enum = append(enum, v.Str())
		}
	}

	return enum
}
