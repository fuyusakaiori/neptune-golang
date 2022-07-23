package znet

import (
	"fmt"
	"neptune-go/src/zinx/utils"
	"neptune-go/src/zinx/ziface"
	"net"
)

type Connection struct {
	// 连接 ID
	ConnID uint32
	// 连接
	Conn *net.TCPConn
	// 连接状态
	isClosed bool
	// 处理器
	Router ziface.IRouter
	// TODO 处理布尔值的通道 (goroutine)
	ExitChan chan bool
}

func (conn *Connection) StartConn() {
	fmt.Println("Conn Start... ConnID", conn.ConnID)
	// 1. 执行读取函数
	go conn.ReadConn()
	// 2. 执行写入函数
}

func (conn *Connection) StopConn() {
	fmt.Println("Conn Stop.. ConnID", conn.ConnID)
	// 1. 检查连接是否已经关闭
	if conn.isClosed {
		return
	}
	// 2. 如果没有关闭, 那么关闭连接
	conn.isClosed = true
	if err := conn.Conn.Close(); err != nil {
		fmt.Println("Conn Stop err, ConnID", err, conn.ConnID)
	}
	// TODO 3. 释放管道资源
	close(conn.ExitChan)
}

func (conn *Connection) GetTCPConn() *net.TCPConn {
	// 注: 不要写成递归调用
	return conn.Conn
}

func (conn *Connection) GetConnID() uint32 {
	return conn.GetConnID()
}

func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

func (conn *Connection) Send(data []byte) error {
	return nil
}

func (conn *Connection) ReadConn() {
	fmt.Println("Reader Goroutine is Running... ConnID", conn.ConnID)
	// 1. 函数退出后释放资源
	defer fmt.Println("Reader Goroutine is Exit... ConnID", conn.ConnID)
	defer conn.Conn.Close()
	// 2. 读取数据
	for {
		buf := make([]byte, utils.Config.ZinxMaxPackage)
		_, err := conn.Conn.Read(buf)
		if err != nil {
			fmt.Println("Reader Goroutine read buf err", err)
			return
		}
		// 3. 封装请求
		req := Request{
			Message: buf,
			Conn:    conn,
		}
		// 4. 处理数据
		go func(request ziface.IRequest) {
			conn.Router.PreHandle(request)
			conn.Router.Handle(request)
			conn.Router.PostHandle(request)
		}(&req)
	}
}

func (conn *Connection) WriteConn() {

}

func NewConn(connID uint32, conn *net.TCPConn, router ziface.IRouter) *Connection {
	connection := &Connection{
		ConnID:   connID,
		Conn:     conn,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return connection
}
