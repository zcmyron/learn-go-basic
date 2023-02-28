package main

import (
	"fmt"
	"io"
	"net"
)

func response() string {
	resp := `HTTP/1.1 200 OK
Server: myserver
Content-Type: text/html

This is the body.`
	return resp
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8099")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listener is created successfully, wait for the clients...")
	//defer lis.Close()

	for {
		client, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//defer client.Close()
		func(c net.Conn) {
			for {
				buf := make([]byte, 4096)
				n, err := c.Read(buf)
				if err != nil {
					if err == io.EOF {
						return
					}
					fmt.Println(err)
					return
				}
				fmt.Printf("Read %d, the content is %s", n, string(buf[:n]))
				c.Write([]byte(response()))
			}
		}(client)
	}

}
