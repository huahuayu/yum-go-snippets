package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()     // Sequence和IntSlice可以转换
	return fmt.Sprint([]int(s)) // Sequence和[]int也可以转换
}

func main() {
	s := Sequence{1, 4, 2, 3}

	fmt.Println(s)
}
