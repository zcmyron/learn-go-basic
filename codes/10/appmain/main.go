package main

import (
	"fmt"
	md "learn-go-basic/models"
)

func main() {
	// var news = NewsModel{1, "test", "test"}
	// news := NewsModel{1, "test", "test1"}
	// news := NewsModel{}
	news := new(md.NewsModel)
	(*news).NewsContent = "1"
	fmt.Println(*news)
	fmt.Println(news.ToJSON())
}
