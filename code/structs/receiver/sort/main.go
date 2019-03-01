package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string { // String方法重定义了输出格式，返回排序后的字符串
	sort.Sort(s) // Sort()方法的入参为一个interface，只要Sequence类型实现了Len，Less，Swap方法就可以调用此方法进行排序
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func main() {
	s := Sequence{4, 2, 1, 3}

	fmt.Println(s)
}
