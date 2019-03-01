package main

import (
	"GoCasts/code/visibility/pkg1"
	"GoCasts/code/visibility/pkg2"
	"fmt"
)

func main() {
	pkg1.InitVar()
	fmt.Println("package1 var from main:", pkg1.Name, pkg1.Age)
	name, age := pkg2.ReadPkg1Var()
	fmt.Printf("package1 var from pkg2:%v,%v", name, age)
}
