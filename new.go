package main

import (
	"fmt"
	"sync"
)

// 对于应用类型的变量，我们不仅需要声明，还需要为其分配内存空间
type User struct {
	lock sync.RWMutex
	name string
	age int
}

func main() {

	i := new(int)
	*i = 10
	fmt.Println(*i)

	//
	u := new(User)
	u.lock.Lock()
	u.name = "张三"
	u.lock.Unlock()
	fmt.Println(u)
}
