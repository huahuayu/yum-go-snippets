package main

import (
	"fmt"
)

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	back1, back2 := -1, 1 // 预先定义好前两个值

	return func() int {
		// 重新赋值(这个就是核心代码)
		back1, back2 = back2, (back1 + back2)
		return back2
	}

}
