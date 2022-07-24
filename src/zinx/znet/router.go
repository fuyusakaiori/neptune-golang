package znet

import (
	"fmt"
	"neptune-go/src/zinx/ziface"
)

type Router struct {
	// 处理器集合
	Apis map[uint32]ziface.IHandler
}

func NewRouter() ziface.IRouter {
	return &Router{
		Apis: make(map[uint32]ziface.IHandler),
	}
}

func (router *Router) RouterHandler(request ziface.IRequest) {
	// 1. 获取处理器: 如果没有找到, 那么返回类型对应的零值; 如果存在, 那么就返回对应值
	handler, result := router.Apis[request.GetMessage().GetMessageID()]
	// 2. 检查是否存在
	if !result {
		fmt.Println("[zinx] router not found handler to handle message")
		return
	}
	// 3. 处理消息
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (router *Router) AddHandler(id uint32, handler ziface.IHandler) {
	// 1. 检查是否存在
	if _, result := router.Apis[id]; result {
		fmt.Println("[zinx] already exit same message id handler in router ")
		return
	}
	// 2. 添加处理器
	router.Apis[id] = handler
}
