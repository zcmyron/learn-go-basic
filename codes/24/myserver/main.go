package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8099")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listener is created successfully, wait for the clients...")
	defer lis.Close()
	client, err := lis.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	for {
		buf := make([]byte, 5)
		n, err := client.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("Read %d, the content is %s", n, string(buf))
	}

}
