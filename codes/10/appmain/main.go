package main

import (
	"fmt"

	// "github.com/zcmyron/learn-go-basic/codes/10/models"
	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"
	core "github.com/zcmyron/learn-go-basic/codes/10/core"
	services "github.com/zcmyron/learn-go-basic/codes/10/servicesb"
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
	// var service services.IService = new(services.NewsService)
	// fmt.Println(service.Get(1))

	// var service services.IService = new(services.ServiceFactory).Create("user")
	var service core.IService = services.NewServiceFactory().Create("news")
	fmt.Println(service.Get(1))

}
