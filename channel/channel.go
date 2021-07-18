package main

import "fmt"

func main() {
	// 定义一个 channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行...")

		c <- 666 // 将 666 发送给 c

	}()
	num := <-c // 从 c 中接受数据，并赋值给 num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
