package main

import (
	. "com.jtthink/core"
	 "fmt"
	_ "com.jtthink/servicesb"
)



func main()  {


	fmt.Println(GetService().Get(123))

}




