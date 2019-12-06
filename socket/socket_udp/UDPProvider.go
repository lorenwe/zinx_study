package main

import (
	"fmt"
	"log"
	"net"
)

// udp服务提供者

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 1234})
	if err != nil {
		log.Println("监听端口错误:", err)
		return
	}
	defer listener.Close()
	log.Println("监听成功")
	//fmt.Printf("地址为: <%s> \n", listener.LocalAddr().String())

	data := make([]byte, 1024)

	for {
		n, remoteAddr, err := listener.ReadFromUDP(data) //读取数据
		if err != nil {
			fmt.Printf("读取失败: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		n, err = listener.WriteToUDP([]byte("world"), remoteAddr) //往发送方，写入数据
		if err != nil {
			fmt.Printf("发送失败: %s", err.Error())
		}
		fmt.Printf("发送消息长度 %d 对方地址: <%s>\n", n, remoteAddr)
	}
}
