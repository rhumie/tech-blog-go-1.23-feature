//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	it := func(yield func(string) bool) {
		for _, s := range []string{"a", "b", "c"} {
			if !yield(s) {
				return
			}
		}
	}

	for i, s := range slices.Collect(it) {
		fmt.Printf("%d: %s\n", i, s)
	}
}
