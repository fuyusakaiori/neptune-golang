package znet

import (
	"neptune-go/src/zinx/ziface"
)

type BaseRouter struct {
}

func (router *BaseRouter) PreHandle(request ziface.IRequest) {
}

func (router *BaseRouter) Handle(request ziface.IRequest) {

}

func (router *BaseRouter) PostHandle(request ziface.IRequest) {

}
