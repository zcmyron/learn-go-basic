package main

import (

	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
	// fmt.Println(core.GetService().Get(1))
	db, err := sql.Open("mysql", "root:ZCmyron27@tcp(127.0.0.1:3306)/test1?charset=utf8mb4")
	if err != nil {
		fmt.Println("connection error: " + err.Error())
		return
	} else {
		fmt.Println("DB connection successful!")
	}
	rows, err := db.Query("select * from test_table_1")
	if err != nil {
		fmt.Println("query error: " + err.Error())
		return
		// panic(err)
	} else {
		fmt.Println("Query successfully!")
	}
	// fmt.Println(rows.Columns())

	// userModel := um.UserModel{}
	// for rows.Next() {
	// 	// var uid int
	// 	// var uname string
	// 	rows.Scan(&userModel.Uid, &userModel.Uname)
	// 	fmt.Println(userModel)
	// }
	// userModels := []um.UserModel{}
	// for rows.Next() {
	// 	temp := um.UserModel{}
	// 	rows.Scan(&temp.Uid, &temp.Uname)
	// 	userModels = append(userModels, temp)
	// }
	// fmt.Println(userModels)
	// allRows := make([]interface{}, 0)
	// for rows.Next() {
	// 	oneRow := make([]interface{}, 2)
	// 	rows.Scan(&oneRow[0], &oneRow[1])
	// 	for i, val := range oneRow {
	// 		v, ok := val.([]byte)
	// 		if ok {
	// 			oneRow[i] = string(v)
	// 		}
	// 	}
	// 	allRows = append(allRows, oneRow)
	// }
	// fmt.Println(allRows)
	columns, _ := rows.Columns()
	allRows := make([]interface{}, 0)
	for rows.Next() {
		fieldMap := make(map[string]string)
		oneRow := make([]interface{}, len(columns))
		scanRow := make([]interface{}, len(columns))
		for i, _ := range oneRow {
			scanRow[i] = &oneRow[i]
		}
		rows.Scan(scanRow...)
		for i, val := range oneRow {
			v, ok := val.([]byte)
			if ok {
				fieldMap[columns[i]] = string(v)
			}
		}
		allRows = append(allRows, fieldMap)

	}
	fmt.Println(allRows...)

}
