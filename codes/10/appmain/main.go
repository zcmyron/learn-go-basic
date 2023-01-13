package main

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
)

type NewsModel struct {
	NewsID                 int
	NewsTitle, NewsContent string
}

func (news NewsModel) ToJSON() string {
	res, err := ffjson.Marshal(news)
	if err != nil {
		return err.Error()
	} else {
		return string(res)
	}
}

func main() {
	// var news = NewsModel{1, "test", "test"}
	// news := NewsModel{1, "test", "test1"}
	// news := NewsModel{}
	news := new(NewsModel)
	(*news).NewsContent = "1"
	fmt.Println(*news)
	fmt.Println(news.ToJSON())
}
