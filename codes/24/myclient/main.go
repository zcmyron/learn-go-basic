package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8099")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("I am smart"))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("read %d from server: %s", n, string(buf[:n]))
}
