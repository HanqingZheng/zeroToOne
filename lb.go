package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type ss struct {
	list  []string
	index int
}

//139.196.94.214
var list = ss{
	list:  []string{"127.0.0.1:8001", "127.0.0.1:8002"},
	index: 0,
}

func main() {
	run()
}

func run() {
	//tcpaddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080");
	lis, err := net.Listen("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("建立连接错误:%v\n", err)
			continue
		}
		fmt.Println(conn.RemoteAddr(), conn.LocalAddr())
		go handle(conn)
	}
}
func handle(sconn net.Conn) {
	fmt.Println(1)
	defer sconn.Close()
	ip := list.list[list.index]
	list.index = (list.index + 1) % len(list.list)

	fmt.Println(ip)

	dconn, err := net.DialTimeout("tcp", ip, 2*time.Second)
	if err != nil {
		fmt.Printf("连接%v失败:%v\n", ip, err)
		return
	}
	go func(sconn net.Conn, dconn net.Conn) {
		_, err := io.Copy(dconn, sconn)
		fmt.Printf("往%v发送数据失败:%v\n", ip, err)
	}(sconn, dconn)
	io.Copy(sconn, dconn)
}

func ReadRequest(lconn net.Conn, rconn net.Conn) {
	io.Copy(rconn, lconn)
}

func ReadResponse(lconn net.Conn, rconn net.Conn) {
	io.Copy(rconn, lconn)
}
