package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f1 int    `tag1:"v1" tag2:"v2"`
	f2 string `one:"v1" two:"v2" blank:""`
}

func main() {
	t := reflect.TypeOf(T{})
	fmt.Println(t.Name())            // 类型名
	fmt.Println(t.FieldByName("f1")) // 字段f1的所有元数据

	f, _ := t.FieldByName("f2")
	fmt.Println(f.Tag) // 字段f2的所有tag

	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 字段f2的 "one" tag

	v, ok = f.Tag.Lookup("two")
	fmt.Printf("%s, %t\n", v, ok)

	v, ok = f.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok)

	v, ok = f.Tag.Lookup("nothing")
	fmt.Printf("%s, %t\n", v, ok)
}
