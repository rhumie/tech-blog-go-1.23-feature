//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	values := []string{"a", "b", "c"}

	for i, s := range slices.Backward(values) {
		fmt.Printf("%d: %s\n", i, s)
	}
}
