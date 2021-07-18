package main

import "fmt"

// interface{} 万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	// interface{} 如何区分 数据类型
	// 类型断言机制
	value,ok := arg.(string)
	if ok {
		fmt.Println("arg is string type")
	}else {
		fmt.Println("arg is not string type")
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}

