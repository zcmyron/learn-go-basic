package main



import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"com.jtthink/models"
)



func main()  {

	db, err := sql.Open("mysql", "shenyi:123123@tcp(localhost:3306)/test?charset=utf8mb4")
	if err != nil{
		fmt.Println("连接错"+err.Error())
		return
	}
	rows,err:=db.Query("select user_id,user_name from users")
	if err != nil{
		fmt.Println("查询错误"+err.Error())
		return
	}
	userModel:=models.UserModel{}
	for rows.Next(){


		rows.Scan(&userModel.Uid,&userModel.Uname)
		fmt.Println(userModel)
	}


	fmt.Println(userModel)


}




