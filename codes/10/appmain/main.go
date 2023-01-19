package main

import (
	"github.com/zcmyron/learn-go-basic/codes/10/models"
	"github.com/zcmyron/learn-go-basic/docs/gobasic/10/课件/models"
)

func main() {
	// var news = NewsModel{1, "test", "test"}
	// news := NewsModel{1, "test", "test1"}
	// news := NewsModel{}
	// news := new(models.NewsModel)
	// (*news).NewsContent = "1"
	// fmt.Println(*news)
	// fmt.Println(news.ToJSON())

	// sportnews := submodels.SportsNews{}
	// fmt.Println(sportnews)
	var service Iservice = models.NewsModel()

}
