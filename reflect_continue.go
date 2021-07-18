package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u *User) Call() {
	fmt.Println("User is called..")
	fmt.Printf("%v\n",u)
}

func (u *User) Print() {
	fmt.Println("another method")
}

func main() {
	user := User{1,"张三", 20}
	user.Call()

	DoFileAndMethod(user)
}

func DoFileAndMethod(input interface{}) {
	//// 获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is: ", inputType.Name())

	//获取 input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is: ", inputValue)

	// 通过 type 获取里面的value
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", filed.Name, filed.Type,value)
	}

	// 通过 type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Name,m.Type)
	}
}