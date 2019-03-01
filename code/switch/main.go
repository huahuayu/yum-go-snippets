package main

import "fmt"

func main() {
	score := 79
	switch score {
	case 60:
		fmt.Println("60")
	}

	var a int
	a = 10

	if 1 == 1 {
		a := 5
		fmt.Println("inside if", a)
	}

	fmt.Println("outside if", a)

}
