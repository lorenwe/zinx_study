package main

import "zinx/znet"

func main() {
	// 创建一个server句柄，使用zinx 的aapi
	s := znet.NewServer("zinxv0.1")
	s.Server()
}
