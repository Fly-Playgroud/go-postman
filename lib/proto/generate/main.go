// Package main ...
package main

func main() {
	schema := getSchema(schemaV21{})
	parse(schema)
}
