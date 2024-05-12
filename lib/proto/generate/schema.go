package main

import (
	"fmt"
	"github.com/ysmood/gson"
)

type schema interface {
	fmt.Stringer
	Version() string
}

type schemaV20 struct{}

func (v schemaV20) String() string {
	return "https://schema.postman.com/collection/json/v2.0.0/draft-07/collection.json"
}

func (v schemaV20) Version() string {
	return parseSchemaVersion(v)
}

type schemaV21 struct{}

func (v schemaV21) String() string {
	return "https://schema.postman.com/collection/json/v2.1.0/draft-07/collection.json"
}

func (v schemaV21) Version() string {
	return parseSchemaVersion(v)
}

type objType int

const (
	objTypeStruct    objType = iota // such as object
	objTypePrimitive                // such as string bool, etc.
)

type property struct {
	name       string
	ref        string
	definition *definition
	global     gson.JSON
}

type definition struct {
	property    *property
	objType     objType
	typeName    string
	enum        []string
	title       string
	description string
	oneOf       []gson.JSON
	props       []*definition
}

func parse(schema gson.JSON) []*property {
	var properties []*property

	for name, property := range schema.Get("properties").Map() {
		properties = append(properties, parseProperty(name, schema, property))
	}
	return properties
}

func parseProperty(name string, schema, gProperty gson.JSON) *property {
	property := &property{
		name:   name,
		global: schema,
	}

	if gProperty.Has("ref") {
		// if a property only has ref, it's def in definitions property
		property.ref = gProperty.Get("ref").Str()

	} else {
		// else if a property doesn't have ref, it's def just in properties level
		def := &definition{
			property: property,

			// notice: if a def has oneOf property, it must be processed deeply,
			// just like an interface.
			oneOf:       gProperty.Get("oneOf").Arr(),
			objType:     objTypePrimitive,
			title:       gProperty.Get("title").Str(),
			typeName:    gProperty.Get("title").Str(),
			enum:        enumList(gProperty.Get("enum")),
			description: gProperty.Get("description").Str(),
		}
		property.definition = def
	}
	return property
}

func refDef(global gson.JSON, ref string) gson.JSON {
	defField := "definitions"
	def := global.Get(defField).Get(ref)
	return def
}
