package main

import (
	"fmt"
)

func main() {

	s := []string{"alice", "bob", "frank", "shiming"}
	//fmt.Println(formatString(s))
	fmt.Println(formatString2(s, 0))
	fmt.Println(formatString2(s, 1))
	fmt.Println(formatString2(s, 2))
	fmt.Println(formatString2(s, 3))
	fmt.Println(formatString2(s, 4))
	fmt.Println(formatString2(s, 5))
}

func formatString(s []string) string {
	var str string

	switch len(s) {
	case 0:
		return " "
	case 1:
		return s[0]
	case 2:
		return s[0] + " and " + s[1]
	default:
		for i := 0; i < len(s); i++ {
			if i == len(s)-1 {
				str = str + " and " + s[i]
				continue
			}

			if i == len(s)-2 {
				str = str + s[i]
				continue
			}

			str += s[i] + ","
		}
		return str
	}
}

func formatString2(s []string, l int) string {
	var str string
	var left int

	if len(s) > l {
		left = len(s) - l
	} else {
		return formatString(s)
	}

	if l == 0 {
		return ""
	}

	for i := 0; i < l; i++ {
		if i == l-1 {
			str += s[i]
		} else {
			str += s[i] + ","
		}
	}

	str = fmt.Sprintf("%s and %d more", str, left)
	return str
}
