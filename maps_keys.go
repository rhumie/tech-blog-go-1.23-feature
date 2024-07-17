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

	for i := range maps.Keys(m) {
		fmt.Printf("%d\n", i)
	}
}
