package main

import (
	"fmt"
	"strings"
	"unique"
	"unsafe"
)

func Canonicalize(ss []string) {
	for i, s := range ss {
		h := unique.Make(s)
		ss[i] = h.Value()
	}
}

func main() {

	ss := []string{
		strings.Repeat("x", 6),
		strings.Repeat("x", 6),
		strings.Repeat("x", 6),
	}

	fmt.Println("Before:")
	for _, s := range ss {
		fmt.Println(unsafe.StringData(s))
	}

	Canonicalize(ss)

	fmt.Println("After:")
	for _, s := range ss {
		fmt.Println(unsafe.StringData(s))
	}
}
