package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var a string
	// pair<type:string, value: "abcd">
	a = "abcd"

	// pair<type:string, value:"abcd">
	var allType interface{}
	allType = a

	value, _ := allType.(string)
	fmt.Println(value)

	fmt.Println("==============")
	// tty:pair<type:*os.File, value:"/dev/tty">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR,0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	//r:pair<type:value:>
	var r io.Reader
	//r:pair<type:*os.File value:"/dev/tty"文件描述符>
	r = tty

	// w:pair<type:*os.File value:"/dev/tty"文件描述符>
	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("hello this is a test!"))



}