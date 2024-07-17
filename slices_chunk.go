//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	values := []string{"a", "b", "c", "d", "e"}

	for s := range slices.Chunk(values, 2) {
		fmt.Printf("%v\n", s)
	}
}
