package main

import (

	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"

	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/zcmyron/learn-go-basic/codes/10/services/users"
)

func sum(min int, max int, c chan int) {
	result := 0
	for i := min; i <= max; i++ {
		result += i
	}
	c <- result
}

func main() {
	num := 1000001
	min := 0
	c := make(chan int)
	start := time.Now()
	go sum(min, num/4, c)
	go sum(num/4+1, num/2, c)
	go sum(num/2+1, num*3/4, c)
	go sum(num*3/4+1, num, c)
	res1, res2, res3, res4 := <-c, <-c, <-c, <-c
	end := time.Now()
	fmt.Println(end.Sub(start), res1+res2+res3+res4)
}
