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
	upd_addr, err := net.ResolveUDPAddr("udp4", ":2001")
	listen, err := net.ListenUDP("udp4", upd_addr)
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	defer listen.Close()
	fmt.Println("listen Start...:")
	//开启一个Goroutine，处理链接
	fmt.Printf("listen Started at %v\n", listen.LocalAddr())
	fmt.Println("Conn Established")

	//s, err := net.ResolveUDPAddr("udp4", ":2001")

	for {
		conn, err := net.Dial("udp4", ":2000") // Jerry
		if err != nil {
			fmt.Printf("dial failed, err:%v\n", err)
			return
		}
		defer conn.Close()

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("输入新消息:")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
			break
		}
		data = strings.TrimSpace(data)
		go func() {

			_, err = conn.Write([]byte(data))
			if err != nil {
				fmt.Printf("write failed, err:%v\n", err)
				return
			}
		}()
		go server(listen)
	}
}

//2.接收客户端的链接
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
