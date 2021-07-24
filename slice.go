package main

import "fmt"

func pirntArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println(index,value)
	}
}



func main() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1,2,3,4}


	for i:=0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index,value := range myArray1 {
		fmt.Println(index,value)
	}



	// 动态数组
	myArray3 := []int{1,2,3,4} // 动态数组

	fmt.Printf("my Array is %T", myArray3)
}
