package shandle

import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type StartGame struct {
	snet.BaseRouter
}

func (S *StartGame) Handle(request isface.IRequest) {

	RoomMaster := &Data.PlayerInfo{
		Roomid:   0,
		Uid:      0,
		Username: "",
	}
	//获得房主信息
	proto.Unmarshal(request.GetData(), RoomMaster)

	//获得当前的房间
	NowRoom := snet.RoomMgr.GetRoom(RoomMaster.Roomid)

	for k, _ := range NowRoom.AllPlayer {
		//snet.ConnMap[uint32(k)].State = snet.PLAY
		//snet.ConnMap[uint32(k)].Count = 1
		snet.ConnMap[uint32(k)].UpdateGameDoing()
	}

	fmt.Println("房主点击了开始游戏,房主ID为：", RoomMaster.Uid, "房间ID为", RoomMaster.Roomid, "给所有人房主发送Start_Game_ACK,让房主预处理一局的信息")
	msg := snet.NewMsgPackage(START_GAME_ACK, request.GetData())

	NowRoom.BroadRoom(msg)
}
