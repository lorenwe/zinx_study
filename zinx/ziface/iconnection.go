package ziface

import "net"

// 定义连接模块的抽象层

type IConnection interface {
	// 启动链接 让当前的连接准备开始工作
	Start()
	// 停止链接 结束当前链接的工作
	Stop()
	// 获取当前链接的绑定 socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前链接模块的ID
	GetConnID() uint32
	// 获取远程客户端的TCP状态，IP port
	RemoteAddr() net.Addr
	// 发送数据
	Send(data []byte) error
}

// 定义一个处理链接业务的方法, 绑定链接方法
type HandleFunc func(*net.TCPConn, []byte, int) error
