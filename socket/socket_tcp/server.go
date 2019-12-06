package main

import (
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("监听端口错误:", err)
		return
	}
	defer server.Close()
	log.Println("监听成功")
	var i int
	for {
		//time.Sleep(time.Second * 10)
		conn, err := server.Accept() //这里不停的去接受新的连接，有新连接就继续往下执行
		if err != nil {
			log.Println("建立连接错误:", err)
			break
		}
		i++
		log.Printf("第 %d: 个新的客户端连接\n", i)
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// read from the connection
		var buf = make([]byte, 1024)
		log.Println("开始读取连接里面的内容")
		//conn.SetReadDeadline(time.Now().Add(time.Microsecond * 10)) //设置10微秒的读取超时时间
		n, err := conn.Read(buf) // 这里会保持与客户端的连接，当客户端主动关闭连接后，此处err会有内容
		if err != nil {
			log.Println("连接读取错误:", err)
			//log.Printf("读取了 %d bytes,  error: %s\n", n, err)
			//if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
			//	continue
			//}
			return
		}
		data := string(buf[:n])
		log.Printf("读取到 %d bytes 的客户端发送的数据内容 %s\n", n, data)
		// 发送数据
		n2, err := conn.Write([]byte(data)) //写入数据
		log.Printf("发送了 %d bytes 的数据，内容是 %s\n", n2, data)
		// 尝试主动关闭连接
		//log.Println("服务器主动关闭连接")
		//conn.Close()
		//break
	}
}
