package main

import (
	"fmt"
	"regexp"
)

func GetRegPath(req string) string {
	req := regexp.MustCompile(`^Get\s(.*?)\sHTTP`)
	if req.
}
func main() {
	str:=`GET /abc.php HTTP/1.1
Host: localhost:8099
Connection: keep-alive
Upgrade-Insecure-Requests: 1`
	fmt.Println(GetRegPath(str))
}
