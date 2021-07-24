package main

import "fmt"

func main() {
	// 声明 myMap 是一种 map 类型， key 是 string,value 是 string
	var myMap map[string]string

	// 分配内存空间
	myMap = make(map[string]string, 10)

	myMap["001"] = "go"

	// ==> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[0] = "c++"

	// ==> 第三种声明方式
	myMap3 := map[string]string{
		"one": "php",
		"two": "c",
	}
	fmt.Println(myMap3)

	cityMap := make(map[string]string)

	// 添加
	cityMap["china"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 删除
	delete(cityMap,"USA")

}
