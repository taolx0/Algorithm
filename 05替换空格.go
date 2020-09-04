package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	var b []rune
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(reflect.TypeOf(s[i]))
	//	if s[i] == ' ' {
	//		b = append(b, 37, 50, 48)
	//	} else {
	//		b = append(b, 23)
	//	}
	//}
	for _, v := range s {
		if v == ' ' {
			b = append(b, 37, 50, 48)
		} else {
			b = append(b, v)
		}
	}
	return string(b)
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
