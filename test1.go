package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var x = 100000000000000
	//fmt.Println(x)
	//fmt.Println(unsafe.Sizeof(x))
	//t(x)
	//
	//var y = 23
	//fmt.Println(y)
	//fmt.Println(unsafe.Sizeof(y))
	//t(y)
	var a int64
	fmt.Println(reflect.TypeOf(a))
}

func t(i interface{}) { //函数t有一个参数i
	switch i.(type) { //多选语句switch
	case string:
		//是字符时做的事情
		fmt.Println("this is a string")
		break
	case int:
		//是整数时做的事情
		fmt.Println("this is a int")
		break
	case bool:
		fmt.Println("this is a bool")
		break
	case int32:
		fmt.Println("this is a int32")
		break
	case float32:
		fmt.Println("this is a float32")
		break
	case float64:
		fmt.Println("this is a float64")
		break
	default:
		fmt.Println("this is a default")
	}
	return
}
