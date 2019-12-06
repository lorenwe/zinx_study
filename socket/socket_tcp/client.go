package main

import (
	"log"
	"net"
)

func main() {
	log.Println("开始建立连接")
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	//conn.SetDeadline(time.Now().Add(time.Second*10)) //10S超时
	if err != nil {
		log.Println("建立连接错误:", err)
		return
	}
	defer conn.Close()
	log.Println("连接成功")
	//time.Sleep(time.Second * 2)

	data := "hello"
	//data := make([]byte, 65536)
	n, err := conn.Write([]byte(data)) //写入数据
	//n, err := conn.Write(data) //发送大数据，模拟服务端读取超时
	log.Printf("发送了 %d bytes 的数据，内容是 %s\n", n, data)

	//之后再不停的接收服务端数据
	for {
		var buf = make([]byte, 10)
		n, err := conn.Read(buf) //这里会保持与服务端的连接，当客户端主动关闭连接后，此处err会有内容
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("读取到 %d bytes 的服务端发送的数据内容 %s\n", n, string(buf[:n]))
	}

}
