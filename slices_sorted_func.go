//go:build ignore

package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	it := func(yield func(Person) bool) {
		for _, s := range []Person{
			{"Gopher", 13},
			{"Alice", 20},
			{"Bob", 24},
			{"Alice", 55},
		} {
			if !yield(s) {
				return
			}
		}
	}

	for i, v := range slices.SortedFunc(it, func(s1, s2 Person) int {
		return strings.Compare(s1.Name, s2.Name)
	}) {
		fmt.Printf("%d: %v\n", i, v)
	}
}
