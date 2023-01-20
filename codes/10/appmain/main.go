package main

import (

	// "github.com/zcmyron/learn-go-basic/codes/10/models"
	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"
	"fmt"

	core "github.com/zcmyron/learn-go-basic/codes/10/core"
	_ "github.com/zcmyron/learn-go-basic/codes/10/services/users"
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
	// var service core.IService = services.NewServiceFactory().Create("news")
	// core.SetService(services.NewServiceFactory().Create("news"))
	fmt.Println(core.GetService().Get(1))

}
