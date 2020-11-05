package shandle

import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type JoinRoom struct {
	snet.BaseRouter
}

//加入房间的hand
func (J *JoinRoom) Handle(request isface.IRequest) {
	RoomClient := snet.RoomMgr
	Playinfo := &Data.PlayerInfo{}

	proto.Unmarshal(request.GetData(), Playinfo)

	Player := snet.PlayerConn{
		Player: snet.Player{
			Coord:    snet.Coord{},
			Id:       Playinfo.Uid,
			UserName: Playinfo.Username,
			RoomId:   Playinfo.Roomid,
		},
		IConnection: request.GetConnection(),
	}
	flag := RoomClient.EnterRoom(Playinfo.Roomid, Player)
	//玩家进入房间
	if flag == true {
		connm := snet.ConnMap[request.GetConnection().GetConnID()]
		connm.UpdateJoinRoom(Playinfo.Roomid)
		fmt.Println("id为:", Playinfo.Uid, "的玩家进入了房间号为：  ", Playinfo.Roomid, "")
		//将玩家信息进行广播

		AllPlayer := &Data.RoomPlayInfo{}

		//获取当前房间
		NowRoomTemp := RoomClient.GetRoom(Playinfo.Roomid)
		//获取当前房间的人数
		AllPlayer.PlayerLen = NowRoomTemp.GetLen()
		//这个是玩家map
		AllPlayMap := NowRoomTemp.GetAllPlay()

		//准备查一下房间内的人的nikename
		SendDataUser := &Data.User_info{}

		//把当前房间的人塞入
		for k, value := range AllPlayMap {
			snet.SDB.Debug().Table("user_info").Where("uid=?",value.GetPlayerId()).Scan(&SendDataUser)
			tempInfo := &Data.PlayerInfo{
				Roomid:   Playinfo.Roomid,
				Uid:      value.GetPlayerId(),
				Username: SendDataUser.User_name,
			}
			changeUserName := AllPlayMap[k]
			changeUserName.SetPlayerUserName(SendDataUser.User_name)
			AllPlayer.AllPlayerInfo = append(AllPlayer.AllPlayerInfo, tempInfo)
		}
		//序列化一下
		buf, _ := proto.Marshal(AllPlayer)
		//打包进去
		fmt.Println("房间号为:", NowRoomTemp.GetRoomID(), "给所有玩家发送了所有玩家的信息:", AllPlayer)
		//给所有人广播当前加入的玩家是谁
		msg := snet.NewMsgPackage(JOIN_ROOM_BORAD, buf)

		NowRoomTemp.BroadRoom(msg)
	}
}