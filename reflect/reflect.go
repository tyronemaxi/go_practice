package main

import (
	"fmt"
	"reflect"
)

func reflectnum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

func main() {
	num := 3.14
	reflectnum(num)
}