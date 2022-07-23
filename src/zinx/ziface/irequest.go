package ziface

// IRequest 请求
type IRequest interface {
	// GetMessage 获取数据
	GetMessage() []byte
	// GetConn 获取连接
	GetConn() IConnection
}
