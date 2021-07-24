package main

import "fmt"

func main() {
	slice0 := []int{1,2,3}

	// 声明 slice 是一个切片，但是没有给slice分配空间
	var slice1 []int
	slice1 = make([]int,3) // 开辟 3 个空间，默认为 0

	// 直接声明，并赋予空间
	slice2 := make([]int,3)
	slice2[0] = 1
	fmt.Println(slice0,slice1,slice2)

	var numbers = make([]int, 3, 5)

	// 向切片追加元素，len = 4 cap = 5
	numbers = append(numbers,1)

	// 继续追加 len = 5 cap = 5
	numbers = append(numbers,2)

	// 继续追加，溢出时，系统会自动扩容，将 cap = 10
	numbers = append(numbers, 5)

	fmt.Println(numbers,len(numbers),cap(numbers))
	// len 表示 左右指针的长度，cap 表示容量
	num := numbers[0:2]

	num1 := make([]int, 3)

	// copy()
	copy(num1, numbers)


}
