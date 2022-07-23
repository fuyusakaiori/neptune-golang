package znet

import (
	"fmt"
	"neptune-go/src/zinx/ziface"
)

type BaseRouter struct {
}

func (router *BaseRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("base router before handle... ", request)
}

func (router *BaseRouter) Handle(request ziface.IRequest) {
	fmt.Println("base router handle... ", request)
}

func (router *BaseRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("base router after handle ", request)
}
