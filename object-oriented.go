package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	level int // 私有属性 go 中 使用大（可以访问）小写（不可以访问）的区别方式来规定方法和属性是否被外部的包访问
}

// 定义类的三个方法
// 设置名
func (h *Hero) SetName(newName string) string  {
	h.Name = newName
	return h.Name
}

// 获取英雄名
func (h Hero) GetName()  string  {
	fmt.Println("the name of hero is: ", h.Name)
	return h.Name
}

// 展示英雄信息
func (h Hero) Show() {
	fmt.Println("the name of hero is: ", h.Name)
	fmt.Println("the address of hero is: ", h.Ad)
	fmt.Println("the Level of hero is: ", h.level)
}

func main() {
	// 创建一个实例对象
	h := Hero{Name: "Superman", Ad: 1, level:1}
	h.Show()
	h.GetName()
	fmt.Println("=====================")
	// 修改英雄名
	h.SetName("iron Man")
	h.Show()
	h.GetName()
}