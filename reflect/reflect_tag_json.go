package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王",2000,10,[]string{"周星驰","张柏芝"}}

	// 编译，结构体 --> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码，json --> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("my_movie = %s\n",myMovie )

}
