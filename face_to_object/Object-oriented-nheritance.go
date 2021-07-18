package main

import (
	"fmt"
)

type Human struct {
	name, sex string
}

func (h Human) Eat() {
	fmt.Println("human is eating")
}

func (h Human) work() {
	fmt.Println("human is working")
}

// 继承
type Superman struct {
	Human
	skill string
}

// 重写&&新增方法
func (s *Superman) fighting() {
	fmt.Println("superman is fighting!")
}

func (s *Superman) work() {
	fmt.Println("superman is working")
}

func main() {
	h := Human{name: "Jelly", sex: "male"}
	h.Eat()
	h.work()
	fmt.Println("====================")
	s := Superman{Human{"ironman", "male"},"money"}
	s.Eat()
	s.work()
	s.fighting()

}