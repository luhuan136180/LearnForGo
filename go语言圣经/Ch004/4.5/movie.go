package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...

}

func main() {
	//将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed :%s", err)
	}
	fmt.Printf("%s\n", data)

	//产生整齐缩进的输出
	data, err = json.MarshalIndent(movies, " ", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//定义一个函数内的结构体
	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshal failed:%s", err)
	}
	fmt.Println(titles)
}
