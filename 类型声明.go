package main

import "fmt"

var e int = 100

var x,y,z int = 100,200,300

var xx,yy = 100,"100"

var (
	vv int = 100
	jj bool = true
)

const length = 10

const (
	// 可以在const{} 添加一个关键字 iota,该值会累加，默认值为 0
	BEIJING = iota * 10
	SHANGHAI
	guangzhou

)

func main() {
	var a int
	a = 1
	var b int = 100
	var c = 100

	d := 100
	fmt.Println(a,b,c,d,e,BEIJING,SHANGHAI)
}
