package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("服务器开始监听")
	// net.Listen("tcp", "127.0.0.1:8888") // ipv4支持的写法
	listen, err := net.Listen("tcp", "0.0.0.0:8888") // ipv4支持的写法, ipv6都支持
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Conn=%v\n", conn)
		}
	}

}
