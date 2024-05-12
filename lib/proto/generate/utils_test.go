package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(parseSchemaVersion(schemaV20{}))
}
