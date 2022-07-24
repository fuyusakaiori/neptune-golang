package znet

import (
	"errors"
	"fmt"
	"io"
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
		fmt.Println("connection already close, ConnID", conn.ConnID)
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

func (conn *Connection) SendMessage(id uint32, data []byte) error {
	// 1. 获取编解码器
	codec := NewCodec()
	// 2. 封装消息
	message := NewMessage(id, data)
	// 3. 编码
	buf, err := codec.Encode(message)
	if err != nil {
		fmt.Println("[zinx] send encode buf err", err)
		return err
	}
	// 4. 发送数据
	if _, err := conn.Conn.Write(buf); err != nil {
		return errors.New("[zinx] send buf err")
	}
	return nil
}

func (conn *Connection) ReadConn() {
	fmt.Println("Reader Goroutine is Running... ConnID", conn.ConnID)
	// 1. 函数退出后释放资源
	defer fmt.Println("Reader Goroutine is Exit... ConnID", conn.ConnID)
	defer conn.Conn.Close()
	for {
		// 2. 获取定长解码器
		codec := NewCodec()
		// 3. 读取消息体的头信
		headBuf := make([]byte, codec.GetHeadLength())
		if _, err := io.ReadFull(conn.Conn, headBuf); err != nil {
			fmt.Println("[zinx] read head buf err", err)
			return
		}
		// 4. 解码器
		message, err := codec.Decode(headBuf)
		// TODO 暂时没有考虑接收到的消息序列号
		if err != nil || message.GetMessageID() < 0 {
			fmt.Println("[zinx] read decode head buf err", err)
			return
		}
		// 5. 读取消息体
		// TODO 暂时没有考虑解决半包问题
		dataBuf := make([]byte, message.GetMessageLength())
		if _, err := io.ReadFull(conn.Conn, dataBuf); err != nil {
			fmt.Println("[zinx] read decode body buf err", err)
			return
		}
		// 6. 向消息体中填充内容
		message.SetMessageData(dataBuf)
		// 7. 封装请求
		req := Request{
			Message: message,
			Conn:    conn,
		}
		// 4. 处理数据
		go conn.Router.RouterHandler(&req)
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
