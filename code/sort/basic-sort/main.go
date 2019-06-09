package main

import (
	"fmt"
	"sort"
)

func main() {
	is := []int{4, 2, 3, 1}
	sort.Ints(is)
	fmt.Println(is) // [1 2 3 4]

	fs := []float64{4.2, 5.23, 1.1, 3.56}
	sort.Float64s(fs)
	fmt.Println(fs)

	ss := []string{"pears", "apple", "banana", "orange"}
	sort.Strings(ss)
	fmt.Println(ss)
}
