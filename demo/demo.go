package main

import (
	"fmt"
)

var variables = map[string]string{"var1": "test"}

const constants = "constants"

type demo struct {
	a string
	b bool
}

func NewDemo() *demo {
	return &demo{a: "test", b: true}
}

func (d *demo) String() string {
	return fmt.Sprintf("%s", d.a)
}
