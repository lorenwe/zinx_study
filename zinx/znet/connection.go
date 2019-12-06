package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// 连接模块
type Connection struct {
	Conn *net.TCPConn
	// 连接ID
	ConnID uint32
	//当前链接状态
	idClosed bool
	//当前链接所绑定的业务处理方法API
	HandleApi ziface.HandleFunc
	//告知当前连接已经退出/停止的 channnel
	EXitChan chan bool
}

// 初始化连接模块
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		idClosed:  false, // 开启
		HandleApi: callback_api,
		EXitChan:  make(chan bool, 1),
	}

	return c
}

//链接的读业务
func (c *Connection) StartReader() {
	fmt.Println("Read Goroutine is running...")
	defer fmt.Println("ConnID = ", c.ConnID, "Reader is exit, remot addr is ", c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 调用当前连接所绑定的HandleApi
		if err := c.HandleApi(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "handle is error", err)
			break
		}
	}
}

// 启动链接 让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start() ...ConnID = ", c.ConnID)
	// 启动从当前链接的读数据的业务
	go c.StartReader()
	// TODO 启动从当前连接写数据的业务

}

// 停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop() ...ConnID = ", c.ConnID)
	if c.idClosed == true {
		return
	}
	c.idClosed = true

	//回收资源
	c.Conn.Close()
	close(c.EXitChan)
}

// 获取当前链接的绑定 socket conn
//func (c *Connection)GetTCPConnection() *net.TCPConn {
//
//}
// 获取当前链接模块的ID
//func (c *Connection)GetConnID() uint32 {
//
//}
// 获取远程客户端的TCP状态，IP port
//func (c *Connection)RemoteAddr() net.Addr {
//
//}
// 发送数据
//func (c *Connection)Send(data []byte) error {
//
//}
