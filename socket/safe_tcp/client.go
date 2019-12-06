package main

import (
	"awesomeProject/socket/safe_tcp/codec"
	"bufio"
	"log"
	"net"
	"strconv"
	"time"
)

func send_mes(conn net.Conn) {

	for i := 0; i < 100; i++ {
		data := "22222222222222222ppiashjlkajbadfhpojojhpadfobjdfs" + strconv.Itoa(i)
		//data := make([]byte, 65536)
		byteData, err := codec.Encode(data)
		if err != nil {
			log.Println("数据处理失败:", err)
		}
		time.Sleep(time.Duration(1) * time.Second)

		conn.Write(byteData) //写入数据
		conn.Write(byteData)
	}

}

func send_mes2(conn net.Conn) {

	for i := 0; i < 100; i++ {
		data := "asfsafasfsafqafqwfasfsafqw22222222222222222ppiashjlkajbadfhpojojhpadfobjdfs" + strconv.Itoa(i)
		//data := make([]byte, 65536)
		byteData, err := codec.Encode(data)
		if err != nil {
			log.Println("数据处理失败:", err)
		}
		time.Sleep(time.Duration(1) * time.Second)

		conn.Write(byteData) //写入数据
		conn.Write(byteData)
	}

}

func main() {
	log.Println("开始建立连接")
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	//conn.SetDeadline(time.Now().Add(time.Second*10)) //10S超时
	if err != nil {
		log.Println("建立连接错误:", err)
		return
	}
	defer conn.Close()
	log.Println("连接成功")
	//time.Sleep(time.Second * 2)

	// 模拟粘包
	go send_mes(conn)
	go send_mes2(conn)

	//log.Printf("发送数据 %s\n", data)

	//n, err := conn.Write(byteData) //写入数据
	//n, err := conn.Write(data) //发送大数据，模拟服务端读取超时
	//log.Printf("发送数据 %s\n", data)

	//之后再不停的接收服务端数据
	for {
		//var buf = make([]byte, 10)
		//n, err := conn.Read(buf) //这里会保持与服务端的连接，当客户端主动关闭连接后，此处err会有内容
		data, err := codec.Decode(bufio.NewReader(conn))
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("读取数据 %s\n", data)
	}

}
