package ziface

import "net"

// IConnection 客户端连接处理器
type IConnection interface {
	// StartConn 建立连接
	StartConn()
	// StopConn 断开连接
	StopConn()
	// GetTCPConn 获取连接
	GetTCPConn() *net.TCPConn
	// GetConnID 获取连接 ID
	GetConnID() uint32
	// RemoteAddr 获取客户端状态: 连接装填、IP 地址、端口号
	RemoteAddr() net.Addr
	// Send 发送数据
	Send(data []byte) error
}

// HandlerFunc 绑定的业务处理函数
type HandlerFunc func(*net.TCPConn, []byte, int) error
