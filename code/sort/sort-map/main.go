package main

import (
	"fmt"
	"sort"
)

func main() {
	// A map is an unordered collection of key-value pairs. If you need a stable iteration order, you must maintain a separate data structure.
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	// with Go 1.12, the fmt package prints maps in key-sorted order to ease testing
	fmt.Println(m)

	// but actually they are not in order
	for k, v := range m {
		fmt.Printf("%s, %d\n", k, v)
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	// Output:
	// Alice 2
	// Bob 3
	// Cecil 1
}
