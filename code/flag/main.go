package main

import (
	"flag"
	"fmt"
)

var (
	name      string
	age       int
	isStudent bool
)

func init() {
	// flag用法：
	// 第一个参数 ：接收flagname的实际值的
	// 第二个参数 ：flag名称为flagname
	// 第三个参数 ：flagname默认值为1234
	// 第四个参数 ：flagname的提示信息
	// 定义完flag之后要调用flag.Parse()进行解析
	// 程序编译完了可以使用 ./flag -h查看帮助
	flag.StringVar(&name, "name", "shiming", "give you name and the program will display it, the default value is shiming")
	flag.IntVar(&age, "age", 20, "give your age, defaule 20")
	flag.BoolVar(&isStudent, "student", true, "")
	flag.Parse()
}
func main() {
	fmt.Printf("you name is %s, age is %d\n", name, age)
	if isStudent {
		fmt.Println("you are a student")
	} else {
		fmt.Println("you are not a student")
	}
}

/*
$ ./flag -name=shiming -age=30 -student=false
you name is shiming, age is 30
you are not a student
*/
