package znet

import (
	"fmt"
	"neptune-go/src/zinx/ziface"
	"net"
)

type Server struct {
	// 服务器名称
	Name string
	// IP 版本
	IPVersion string
	// IP 地址
	IP string
	// 端口号
	Port int
}

// Start 在方法名前声明接受者的方法, 是属于结构体方法
func (server *Server) Start() {
	// 最外层添加异步处理, 避免同步阻塞建立连接
	go func() {
		// 服务器正式启动
		fmt.Printf("[Start] Server Listener at IP :%s, Port :%d\n", server.IP, server.Port)
		// 1. 获取 TCP 对象
		addr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
		// 错误处理
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		// 2. 获取监听器对象
		listener, err := net.ListenTCP(server.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", server.IPVersion, " err", err)
			return
		}

		fmt.Println("start Zinx server success", server.Name, " success, Listening...")
		// 3. 阻塞等待客户端的连接
		for {
			connection, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 4. 处理业务逻辑: go 声明方法异步执行 协程
			go func() {
				for {
					// 4.1 读取客户端连接发送的数据
					buf := make([]byte, 512)
					length, err := connection.Read(buf)
					if err != nil {
						fmt.Println("receive buf err")
						return
					}
					// 4.2 写回客户端发送的数据: _ 可以表示不使用的返回值
					if _, err := connection.Write(buf[0:length]); err != nil {
						fmt.Println("write back buf err", err)
						return
					}
				}
			}()
		}
	}()

}

func (server *Server) Serve() {
	// 1. 启动服务器
	server.Start()

	// TODO 服务器启动后的额外状态

	// 2. 阻塞服务器, 避免主进程结束导致整个服务器停止
	select {}
}

func (server *Server) Stop() {
	// TODO 服务器关闭前释放相应的资源
}

// NewServer 1. 返回值是 IServer 2. 在方法名前没有声明接受者的, 属于公共的方法
func NewServer(name string) ziface.IServer {
	// 变量的声明
	server := &Server{
		Name:      name,
		IP:        "0.0.0.0",
		IPVersion: "tcp4",
		Port:      8999,
	}
	// 接口方法的入参是指针类型, 就需要传入地址, 所以对象需要取址
	return server
}