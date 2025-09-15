package main

import (
	"fmt"
	"reflect"
)

func typeName(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		t := reflect.TypeOf(val)
		if t.Kind() == reflect.Chan {
			fmt.Printf("chan %s\n", t.Elem())
		} else {
			fmt.Println("unknown type")
		}
	}
}

func main() {
	typeName(3)
	typeName("ame")
	typeName(false)

	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	typeName(ch1)
	typeName(ch2)
	typeName(ch3)

	// unknown type
	typeName(3.22)
}
