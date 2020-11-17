package shandle

/*
import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type Chatting struct {
	snet.BaseRouter
}

//加入房间的hand
func (M *Chatting) Handle(request isface.IRequest) {

	chat := &Data.RoomChat{}
	proto.Unmarshal(request.GetData(), chat)

	nowroom := snet.RoomMgr.GetRoom(chat.Roomid)

	fmt.Println("---这个房间 ", chat.Roomid, "的", request.GetConnection().GetConnID(), "发送了", chat.Chatdata)

	msg := snet.NewMsgPackage(CHAT_ROOM_BORAD, request.GetData())

	nowroom.BroadRoom(msg)
}
*/