package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("开始建立连接")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello word"))
		if err != nil {
			fmt.Println("write conn error:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read buf error:", err)
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n", buf[:cnt], cnt)

		time.Sleep(1 * time.Second)
	}
}
