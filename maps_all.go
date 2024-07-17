//go:build ignore

package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[int]string{
		1: "one",
		10: "ten",
		100: "hundred",
	}

	for i, s := range maps.All(m) {
		fmt.Printf("%d: %s\n", i, s)
	}
}
