package main

import (
	"fmt"
	"net"
)

func main() {
	//sip := net.ParseIP("127.0.0.1")// 单播指定地址
	sip := net.ParseIP("192.168.1.255") //广播地址

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 4321}
	dstAddr := &net.UDPAddr{IP: sip, Port: 1234}

	//单播
	//conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	//if err != nil {
	//	log.Println("链接UDP出错:", err)
	//	return
	//}
	//defer conn.Close()

	// 单播发送数据方式
	//n, err := conn.Write([]byte("hello"))
	//if err != nil {
	//	fmt.Printf("发送失败: %s", err)
	//}
	//fmt.Printf("发送消息长度 %d 对方地址: <%s>\n", n, conn.RemoteAddr())

	// 单播接收消息
	//data := make([]byte, 1024)
	//n, remoteAddr, err := conn.ReadFromUDP(data) //读取数据
	//if err != nil {
	//	fmt.Printf("读取失败: %s", err)
	//}
	//fmt.Printf("对方返回消息: <%s> %s\n", remoteAddr, data[:n])

	//广播
	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		fmt.Println(err)
	}

	//广播发送数据方式
	n, err := conn.WriteToUDP([]byte("hello"), dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("发送消息长度 %d \n", n)

	// 广播接收消息
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFrom(data) //读取数据
	if err != nil {
		fmt.Printf("读取失败: %s", err)
	}
	fmt.Printf("对方返回消息: <%s> %s\n", remoteAddr, data[:n])
	//fmt.Printf("对方返回消息: <%s> %s\n", conn.RemoteAddr(), data[:n])
}
