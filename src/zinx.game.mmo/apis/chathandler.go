package apis

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"neptune-go/src/zinx.game.mmo/core"
	"neptune-go/src/zinx.game.mmo/protobuf/pb"
	"neptune-go/src/zinx/ziface"
	"neptune-go/src/zinx/znet"
)

const (
	Pid = "pid"
)

type WorldChatHandler struct {
	znet.BaseHandler
}

func (handler *WorldChatHandler) PreHandle(request ziface.IRequest) {
}

func (handler *WorldChatHandler) Handle(request ziface.IRequest) {
	// 1. 接收消息并解析
	message := &pb.Talk{}
	err := proto.Unmarshal(request.GetMessage().GetMessageData(), message)
	if err != nil {
		fmt.Println("deserialize message err", err)
		return
	}
	// 2. 获取发送消息的玩家 ID
	pid := request.GetConn().GetConnectionProperty(Pid)
	// 3. 获取玩家对象: 断言就是整型
	player := core.WorldObject.GetPlayerByPid(pid.(int))
	// 4. 广播给所有玩家
	player.Talk(message.Content)
}

func (handler *WorldChatHandler) PostHandle(request ziface.IRequest) {

}
