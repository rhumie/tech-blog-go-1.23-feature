//go:build ignore

package main

import (
	"fmt"
	"maps"
)

func main() {

	m := map[int]string{
		1:   "one",
		10:  "ten",
		100: "hundred",
	}

	it := func(yield func(int, string) bool) {
		for i, s := range map[int]string{
			10: "zehn",
			11: "elf",
		} {
			if !yield(i, s) {
				return
			}
		}
	}

	maps.Insert(m, it)

	for i, s := range m {
		fmt.Printf("%d: %s\n", i, s)
	}
}
