package shandle

import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

//退出房间时候的Handle

type ExitRoom struct {
	snet.BaseRouter
}

//需要传入退出房间的PlayInfo
func (E ExitRoom) Handle(request isface.IRequest) {
	PlayerMsg := &Data.PlayerInfo{
		Roomid:   0,
		Uid:      0,
		Username: "",
	}

	proto.Unmarshal(request.GetData(), PlayerMsg)

	SenData := &Data.RoomPlayInfo{
		AllPlayerInfo: nil,
		PlayerLen:     0,
	}
	//如果是房主那么让所有人退出房间。

	NowRoom := snet.RoomMgr.GetRoom(PlayerMsg.Roomid)

	fmt.Println("房间ID为",PlayerMsg.Roomid,"用户ID为",PlayerMsg.Uid)

	if PlayerMsg.Roomid == PlayerMsg.Uid {
		fmt.Println("房主退出房间")

		SenBuf, _ := proto.Marshal(SenData)
		//如果房主退出，那么Len = 0 ，那么房间解散。
		//客户端判断Len -->PlayerLen 是否等于0即可
		SenMes := snet.NewMsgPackage(EXIT_ROOM_ACK, SenBuf)

		NowRoom.BroadRoom(SenMes)
		//删除房间
		for k, v := range NowRoom.AllPlayer {

			snet.ConnMap[uint32(k)].UpdateExitRoom()
			//玩家退出房间
			fmt.Println("玩家",v.Id,"退出房间:",NowRoom.Id)

			v.IConnection.SendMesg(EXIT_ROOM_ACK,[]byte("1"))
		}

		delete(snet.RoomMgr.AllRoom, PlayerMsg.Roomid)

	}else {
		//说明是别的玩家退出游戏，那么就广播当前所有玩家的信息。
		//删除该玩家在房间的信息
		fmt.Println("玩家退出房间")
		delete(NowRoom.AllPlayer, PlayerMsg.Uid)
		//更改链接状态
		snet.ConnMap[uint32(PlayerMsg.Uid)].UpdateExitRoom()
		SenData.PlayerLen = NowRoom.GetLen()
		AllPlayMap := NowRoom.GetAllPlay()
		for _, value := range AllPlayMap {
			tempInfo := &Data.PlayerInfo{
				Roomid:   PlayerMsg.Roomid,
				Uid:      value.GetPlayerId(),
				Username: value.UserName,
			}
			SenData.AllPlayerInfo = append(SenData.AllPlayerInfo, tempInfo)
		}
		fmt.Println("房间为",SenData.AllPlayerInfo[0].Roomid,"发送数据为",SenData)
		SenBuf, _ := proto.Marshal(SenData)

		msg := snet.NewMsgPackage(JOIN_ROOM_BORAD, SenBuf)

		request.GetConnection().SendMesg(EXIT_ROOM_ACK,[]byte("1"))
		NowRoom.BroadRoom(msg)
	}

}
