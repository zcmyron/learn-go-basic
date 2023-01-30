package main

import (

	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"

	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/zcmyron/learn-go-basic/codes/10/services/users"
)

func sum(max int) {
	result := 0
	for i := 0; i <= max; i++ {
		result += i
	}
	fmt.Println(result)
}

func main() {
	defer time.Sleep(time.Second * 3)
	go sum(100)
}
