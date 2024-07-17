//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	it := func(yield func(string) bool) {
		for _, s := range []string{"b", "c", "a"} {
			if !yield(s) {
				return
			}
		}
	}

	for i, s := range slices.Sorted(it) {
		fmt.Printf("%d: %s\n", i, s)
	}
}
