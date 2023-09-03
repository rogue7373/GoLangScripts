package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var l = fmt.Printf

func main() {

	p := point{1, 2}
	l("struct1: %v\n", p)

	l("struct2: %+v\n", p)

	l("struct3: %#v\n", p)

	l("type: %T\n", p)

	l("bool: %t\n", true)

	l("int: %d\n", 123)

	l("bin: %b\n", 14)

	l("char: %c\n", 33)

	l("hex: %x\n", 456)

	l("float1: %f\n", 78.9)

	l("float2: %e\n", 123400000.0)
	l("float3: %E\n", 123400000.0)

	l("str1: %s\n", "\"string\"")

	l("str2: %q\n", "\"string\"")

	l("str3: %x\n", "hex this")

	l("pointer: %p\n", &p)

	l("width1: |%6d|%6d|\n", 12, 345)

	l("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	l("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	l("width4: |%6s|%6s|\n", "foo", "b")

	l("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
