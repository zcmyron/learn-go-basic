package main

import (

	// "github.com/zcmyron/learn-go-basic/codes/10/submodels"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/zcmyron/learn-go-basic/codes/10/services/users"
)

func test() (ret int) {
	defer func() {
		ret = -1
	}()
	panic("exception...")
}

func main() {
	url := "https://www.cnblogs.com/#p%d"

	c := make(chan map[int][]byte)

	for i := 0; i <= 3; i++ {
		go func(index int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}

				// if index == 3 {
				// 	close(c)
				// }
			}()
			url := fmt.Sprintf(url, index)
			res, _ := http.Get(url)
			cnt, _ := ioutil.ReadAll(res.Body)
			c <- map[int][]byte{index: cnt}
		}(i)

	}
	// i := 0
	// for getcnt := range c {
	// 	for k, v := range getcnt {
	// 		ioutil.WriteFile(fmt.Sprintf("../files/%d", k), v, 0777)
	// 	}
	// 	i++
	// 	if i == 3 {
	// 		close(c)
	// 	}
	// }

	result := map[int][]byte{}
myloop:
	for {
		select {
		case result = <-c:
			for k, v := range result {
				ioutil.WriteFile(fmt.Sprintf("../files/%d", k), v, 0777)
			}
		case <-time.After(time.Second * 3):
			break myloop
		}
	}

	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// fmt.Println(test())

}
