package main

import (
	"fmt"

	// "github.com/zcmyron/learn-go-basic/codes/10/models"
	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"
	"github.com/zcmyron/learn-go-basic/codes/10/services"
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
	var service services.IService = services.NewsService
	fmt.Println(service.Get(1))

}
