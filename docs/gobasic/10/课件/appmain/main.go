package main

import (
	"fmt"
	"com.jtthink/services"
)

func main()  {

   var service services.IService=services.NewServiceFactory().Create("user")

	fmt.Print(service.Get(123))

 


}


