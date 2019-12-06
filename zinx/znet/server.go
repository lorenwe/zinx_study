package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务绑定IP
	IP string
	// 服务器监听的端口
	Port int
}

// 定义当前客户端连接所绑定的handle api
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 客户端已经与客户端建立连接
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
	//go func(conn *net.TCPConn) {
	//	for  {
	//		buf := make([]byte, 512)
	//		cnt, err := conn.Read(buf)
	//		if err != nil {
	//			fmt.Println("conn.Read err ", err)
	//			continue
	//		}
	//		// 回显
	//		if _, err := conn.Write(buf[:cnt]); err != nil {
	//			fmt.Println("write back buf err ", err)
	//			continue
	//		}
	//
	//	}
	//
	//}(conn)
}

// 实现IServer的接口
func (s *Server) Start() {
	fmt.Printf("[Start] Server listenner at IP : %s, Port %d is starting\n", s.IP, s.Port)

	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error:", err)
			return
		}
		// 2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err: ", err)
			return
		}
		fmt.Println("start Zinx server succ, ", s.Name, "succ, Listenning...")

		var cid uint32
		// 4 阻塞的等待客户端连接
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP err ", err)
				continue
			}

			// 将处理新连接的业务方法 和conn 进行绑定， 得到连接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			go dealConn.Start()

		}
	}()
}

func (s *Server) Stop() {
	// TODO 将服务器资源状态以及链接停止
}

func (s *Server) Server() {
	s.Start()

	// TODO 做一些启动服务后的额外业务

	// 阻塞状态，防止协程退出
	select {}
}

// 初始化server的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
