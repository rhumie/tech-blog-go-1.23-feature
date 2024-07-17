//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	values := []string{"a", "b", "c"}

	it := func(yield func(string) bool) {
		for _, s := range []string{"d", "e", "f"} {
			if !yield(s) {
				return
			}
		}
	}

	for i, s := range slices.AppendSeq(values, it) {
		fmt.Printf("%d: %s\n", i, s)
	}
}
