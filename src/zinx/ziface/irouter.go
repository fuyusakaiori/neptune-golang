package ziface

type IRouter interface {
	RouterHandler(request IRequest)

	AddHandler(id uint32, handler IHandler)
}
