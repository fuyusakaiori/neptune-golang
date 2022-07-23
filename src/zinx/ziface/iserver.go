package ziface

type IServer interface {
	// Start 启动服务器
	Start()
	// Serve 运行服务器
	Serve()
	// Stop 停止服务器
	Stop()
	// AddRouter 添加处理器
	AddRouter(router IRouter)
}
