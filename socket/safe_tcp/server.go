package main

import (
	"awesomeProject/socket/safe_tcp/mybuffer"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8889")
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
		go doConn(conn)
	}
}

//A.先接收到data1,然后接收到data2.
//
//B.先接收到data1的部分数据,然后接收到data1余下的部分以及data2的全部.
//
//C.先接收到了data1的全部数据和data2的部分数据,然后接收到了data2的余下的数据.
//
//D.一次性接收到了data1和data2的全部数据.

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// read from the connection
		var buf = make([]byte, 1024)
		log.Println("开始读取连接里面的内容")
		n, err := conn.Read(buf) // 这里会保持与客户端的连接，当客户端主动关闭连接后，此处err会有内容
		if err != nil {
			log.Println("连接读取错误:", err)
			return
		}
		lengthByte := buf[:8]
		// 转成int
		bytesBuffer := bytes.NewBuffer(lengthByte)
		var x int32
		binary.Read(bytesBuffer, binary.BigEndian, &x)

		//data := string(buf[8:n]) //把包头内容截掉了
		data2 := string(buf[:n])
		// 读取到完整数据
		fmt.Println(buf)
		log.Printf("头信息是 %s", lengthByte)
		fmt.Println(lengthByte)
		log.Printf("头数据是 %d\n", x)
		log.Printf("数据包总长度是 %d\n", n)
		//log.Printf("读取到客户端发送的数据内容 %s\n", data)

		// 把内容发送出去
		n2, err := conn.Write([]byte(data2))
		if err != nil {
			log.Println("发送数据错误:", err)
			return
		}
		log.Printf("发送数据长度 %d\n", n2)
	}
}

func doConn(conn net.Conn) {
	defer conn.Close()
	var (
		buffer           = bytes.NewBuffer(make([]byte, 0, 65535)) //buffer用来缓存读取到的数据
		readBytes        = make([]byte, 65535)                     //readBytes用来接收每次读取的数据，每次读取完成之后将readBytes添加到buffer中
		isHead           = true                                    //用来标识当前的状态：正在处理size部分还是body部分
		bodyLen   uint32 = 0                                       //表示body的长度
		head             = make([]byte, 8)
	)

	for {
		//首先读取数据
		readByteNum, err := conn.Read(readBytes)
		if err != nil {
			log.Println("连接读取错误:", err)
			return
		}

		//log.Printf("内容是 %s", string(readBytes))

		buffer.Write(readBytes[0:readByteNum]) //将读取到的数据放到buffer中

		// 然后处理数据
		for {
			log.Printf("buffer长度是 %d", uint32(buffer.Len()))
			if isHead {
				if uint32(buffer.Len()) >= 8 {
					isHead = false
					//head := make([]byte, 8)
					_, err = buffer.Read(head) // 每次读取buffer内容都会相应的减少读取的部分数据
					log.Printf("buffer长度是 %d", uint32(buffer.Len()))
					if err != nil {
						log.Fatal(err)
						return
					}
					// 转成uint32
					bytesBuffer := bytes.NewBuffer(head)
					binary.Read(bytesBuffer, binary.BigEndian, &bodyLen)
					// 减去包头长度就是实际内容的总长度
					bodyLen = bodyLen - 8
				} else {
					break
				}
			}

			if !isHead {
				log.Printf("buffer长度是 %d", uint32(buffer.Len()))
				if uint32(buffer.Len()) >= bodyLen {
					body := make([]byte, bodyLen)
					_, err = buffer.Read(body[:bodyLen])
					if err != nil {
						log.Fatal(err)
						return
					}
					// 已经获取到完整内容
					//fmt.Println("received body: " + string(body[:bodyLen]))
					// 将包头信息合并到一起发出去
					msg := make([]byte, len(head)+len(body))
					copy(msg, head)
					copy(msg[len(head):], body)
					// 读取完毕恢复初始信息
					isHead = true
					head = make([]byte, 8)

					// 把内容发送出去
					n2, err := conn.Write(msg)
					if err != nil {
						log.Println("发送数据错误:", err)
						return
					}
					log.Printf("发送数据长度 %d\n", n2)
				} else {
					break
				}
			}
		}
	}
}

func doConn2(conn net.Conn) {
	var (
		buffer      = mybuffer.NewBuffer(conn, 16)
		headBuf     []byte
		contentSize int
		contentBuf  []byte
	)
	for {
		_, err := buffer.ReadFromReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			headBuf, err = buffer.Seek(8)
			if err != nil {
				break
			}
			contentSize = int(binary.BigEndian.Uint16(headBuf))
			if buffer.Len() >= contentSize-8 {
				contentBuf = buffer.Read(8, contentSize)
				fmt.Println(string(contentBuf))
				continue
			}
			break
		}
	}
}
