package main

import "fmt"

// interface 本身是一个指针
type Animal interface {
	Sleep()
	Getcolor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

type Dog struct {
	color string
}

// cat 方法
func (a *Cat) Sleep() {
	fmt.Println("cat is Sleeping")
}

func (c *Cat) Getcolor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "cat"
}

// cat 方法
func (a *Dog) Sleep() {
	fmt.Println("cat is Sleeping")
}

func (c *Dog) Getcolor() string {
	return c.color
}

func (c *Dog) Gettype() string {
	return "dog"
}

func showAnimal(a Animal) {
	a.Sleep()
	fmt.Println("color = ", a.Getcolor())
	fmt.Println("kind = ", a.GetType())
}
func main() {
	var animal Animal
	animal = &Cat{"green"}
	animal.Sleep()
	fmt.Println("=============")

	cat := Cat{"yellow"}
	showAnimal(&cat)

	dog := Cat{"black"}
	showAnimal(&dog)
}