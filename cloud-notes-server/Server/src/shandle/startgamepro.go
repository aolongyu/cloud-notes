package shandle

import (
	Data "data"
	"fmt"
	"github.com/golang/protobuf/proto"
	"isface"
	"snet"
)

type StartGamePro struct {
	snet.BaseRouter
}

//由房主预处理信息给服务器，并且发送到服务器上。
func (S *StartGamePro) Handle(request isface.IRequest) {
	fmt.Println("准备预处理阶段")
	PreLoadData := &Data.PreloadPro{}

	proto.Unmarshal(request.GetData(), PreLoadData)
	fmt.Println("反序列化了")
	AllPlayData := PreLoadData.AllPlayer
	AllResource := PreLoadData.AllMapResource
	Romid := AllPlayData[0].PlayerMsg.Roomid
	//获取房间Id
	NowRoom := snet.RoomMgr.GetRoom(Romid)

	//对房间写数据的时候先上锁，防止别人写数据进去。
	NowRoom.Lock()
	defer NowRoom.Unlock()

	LenPlayData := len(AllPlayData)
	LenResource := len(AllResource)
	//发送的数据
	//SendPlayData := make([]*Data.PlayerPro, LenPlayData)
	SendPlayData := make([]*Data.PlayerPro, 0, LenPlayData)
	//把客户端预处理好的X,Y载入到服务器AllPlayer中
	for i := 0; i < LenPlayData; i++ {
		temp := NowRoom.AllPlayer[AllPlayData[i].PlayerMsg.Uid]
		temp.SetCoord(AllPlayData[i].PlayerCoord.X, AllPlayData[i].PlayerCoord.Y)
	}
	//把客户端预处理号的资源，载入到服务器AllResource中
	for i := 0; i < LenResource; i++ {
		Resource := snet.Coord{
			X: AllResource[i].MapResourceCoord.X,
			Y: AllResource[i].MapResourceCoord.Y,
		}
		value := snet.MapResource{
			Coord: Resource,
			Id:    AllResource[i].Id,
		}
		NowRoom.AllResource[Resource] = value
	}
	fmt.Println("房间ID为", NowRoom.GetRoomID(), "房主预处理的消息:", PreLoadData, "\n")
	tmp := &Data.Account_info{}
	tmp2 := &Data.Hero{}

	//接下来操作数据库，把每个信息读出来。
	for k, v := range NowRoom.AllPlayer {
		snet.SDB.Table("account_info").Where("UID=?", k).Scan(&tmp)

		snet.SDB.Table("hero").Where("UID=? AND hero_id=?", k, tmp.Choosen_hero_id).Scan(&tmp2)

		v.SetPlayerMsg(tmp.Choosen_hero_id, tmp2.Hero_skill, tmp2.Hero_lv, tmp.Weapon_id)
		NowK := NowRoom.AllPlayer[k]
		TempValue := NowK.GetPlayAllMsg()
		SendPlayData = append(SendPlayData, TempValue)
	}
	fmt.Println("房间号为", NowRoom.GetRoomID(), "从数据库中更新了数据，数据为：", SendPlayData)
	//准备要发送的数据

	SendData := &Data.PreloadPro{
		AllPlayer:      SendPlayData,
		AllMapResource: PreLoadData.AllMapResource,
	}

	bufSendData, _ := proto.Marshal(SendData)
	msgSendData := snet.NewMsgPackage(GAME_START_ACK, bufSendData)

	//此处初始化排行榜
	for _, v := range NowRoom.AllPlayer {
		SendDataUser := &Data.User_info{}
		snet.SDB.Debug().Table("user_info").Where("uid=?", v.GetPlayerId()).Scan(&SendDataUser)
		temp := &Data.PlayerInfo{
			Roomid:   100,
			Uid:      v.Id,
			Username: SendDataUser.User_name,
		}
		NowRoom.Ranking.AllPlayerInfo = append(NowRoom.Ranking.AllPlayerInfo, temp)
	}

	NowRoom.Ranking.PlayerLen = NowRoom.GetLen()
	fmt.Println("排行榜设置完毕：", NowRoom.Ranking)
	fmt.Println("房间号为", NowRoom.GetRoomID(), "广播所有玩家数据，数据为：", SendData)
	NowRoom.BroadRoom(msgSendData)
	//time.Sleep(1000 * time.Millisecond)
	NowRoom.ChangeChan <- snet.GAME_STATUS_BATTLE
	//NowRoom.SetValue(snet.GAME_STATUS_BATTLE)
}
