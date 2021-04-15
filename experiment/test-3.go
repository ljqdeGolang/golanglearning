package main

import (
	"fmt"
	"net"
	"time"
)

func test3() {
	timeout := time.Duration(10 * time.Second)
	conn, err := net.DialTimeout("tcp","180.76.76.76:80", timeout)
	if err != nil {
   		fmt.Println("Site unreachable, error: ", err)
	}
	fmt.Println(conn.RemoteAddr().Network(),conn.RemoteAddr().String())

	// conn,err :=net.Dial("tcp","baidu.com:http")
	// if err != nil {
	// 	fmt.Println("Site unreachable, error: ", err)
	// }
	// fmt.Println(conn.RemoteAddr().Network(),conn.RemoteAddr().String())
}