//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	values := []string{"a", "b", "c"}

	for s := range slices.Values(values) {
		fmt.Printf("%s\n", s)
	}
}
