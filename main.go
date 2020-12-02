package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 2001 Jerry 2000 Tom

func main() {
	//1.建立监听端口
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 2000, // Tom
	})
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	fmt.Println("listen Start...:")
	//开启一个Goroutine，处理链接
	go server(listen)
	fmt.Println("listen Started")

	conn, err := net.Dial("udp", ":2000") // Jerry
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}
	fmt.Println("Conn Established...:")
	fmt.Println("输入新消息")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("client Started...:")
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
			break
		}

		data = strings.TrimSpace(data)
		//传输数据到服务端
		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			break
		}
	}

}

func server(listen *net.UDPConn) {
	for {
		//2.接收客户端的链接
		data := make([]byte, 1024)
		_, rAdder, err := listen.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("read failed, err:%v\n", err)
			continue
		} else {
			strData := string(data)
			fmt.Printf("Received from:%v\t mesg: %s\n", rAdder, strData)
		}
	}
}
