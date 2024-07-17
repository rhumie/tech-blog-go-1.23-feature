//go:build ignore

package main

import (
	"fmt"
	"maps"
)

func main() {
	it := func(yield func(int, string) bool) {
		for i, s := range map[int]string{
			1:   "one",
			10:  "ten",
			100: "hundred",
		} {
			if !yield(i, s) {
				return
			}
		}
	}

	for i, s := range maps.Collect(it) {
		fmt.Printf("%d: %s\n", i, s)
	}
}
