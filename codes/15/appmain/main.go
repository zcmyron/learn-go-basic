package main

import (

	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"

	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/zcmyron/learn-go-basic/codes/10/services/users"
)

// func sum(min int, max int, c chan int) {
// 	result := 0
// 	for i := min; i <= max; i++ {
// 		result += i
// 	}
// 	c <- result
// }

func main() {
	// num := 1000001
	// min := 0
	// c := make(chan int)
	// start := time.Now()
	// go sum(min, num/4, c)
	// go sum(num/4+1, num/2, c)
	// go sum(num/2+1, num*3/4, c)
	// go sum(num*3/4+1, num, c)
	// res1, res2, res3, res4 := <-c, <-c, <-c, <-c
	// end := time.Now()
	// fmt.Println(end.Sub(start), res1+res2+res3+res4)
	users := strings.Split("aa,bb,cc,dd,ee", ",")
	ages := strings.Split("1,2,3,4,5", ",")

	c1, c2 := make(chan bool), make(chan bool)
	ret := make([]string, 0)

	go func() {
		for _, v := range users {
			<-c1
			ret = append(ret, v)
			c2 <- true
		}
	}()
	go func() {
		for _, v := range ages {
			<-c2
			ret = append(ret, v)
			c1 <- true
		}

	}()

	c1 <- true

	for v := range ret {
		fmt.Println(v)
	}
}
