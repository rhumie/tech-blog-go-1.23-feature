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

	for s := range maps.Values(m) {
		fmt.Printf("%s\n", s)
	}
}
