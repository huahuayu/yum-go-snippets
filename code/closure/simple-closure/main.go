package main

import "fmt"

/*
闭包是由函数以及创建该函数的词法环境组合而成。这个环境包含了这个闭包创建时所能访问的所有局部变量。
将闭包最清楚的文章： https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Closures
*/

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ") // 3. 而且匿名函数外的变量作为闭包的词法环境没有被释放
	fmt.Print(f(30))
}

func Adder() func(int) int {
	var x int // 1. x是匿名函数外的变量
	return func(delta int) int {
		x += delta // 2. 在匿名函数内却能访问
		return x
	}
}
