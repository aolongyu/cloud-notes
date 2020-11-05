package shandle

import (
	Data "data"
	"fmt"
	"github.com/golang/protobuf/proto"
	"isface"
	"snet"
)

type UpdataRank struct{
	snet.BaseRouter
}

func (u *UpdataRank) Handle(request isface.IRequest){
	recvData := &Data.PlayerInfo{
		Roomid:   0,//杀人的人
		Uid:      0,//被杀的人
		Username: "",
	}
	proto.Unmarshal(request.GetData(),recvData)
	nowRoomId := snet.ConnMap[ uint32(recvData.Uid) ].Roomid
	NowRoom := snet.RoomMgr.GetRoom(nowRoomId)

	fmt.Println(recvData.Roomid,"杀了",recvData.Uid,"，在房间",nowRoomId)

	Rank := NowRoom.Ranking

	NowLen := Rank.PlayerLen
	var i,j int32
	j = -1
	var gold int32
	for i = 0 ;i<NowLen;i++{
		//玩家扣分
		if Rank.AllPlayerInfo[i].Uid == recvData.Uid {
			//增加或减少积分
			//Rank.AllPlayerInfo[i].Roomid = recvData.Roomid
			//玩家积分除2
			Rank.AllPlayerInfo[i].Roomid /= 2
			//需要加分的玩家应该获得的分
			gold = Rank.AllPlayerInfo[i].Roomid
			if(gold < 100){
				Rank.AllPlayerInfo[i].Roomid = 100
			}
		}else if(Rank.AllPlayerInfo[i].Uid == recvData.Roomid){
			j = i
		}
	}
	if j!= -1 {
		Rank.AllPlayerInfo[j].Roomid += gold
	}
	var flag bool
	//进行排序
	for i = 0;i<NowLen;i++{
		flag = true
		for j = 0;j<NowLen-1;j++{
			if Rank.AllPlayerInfo[j].Roomid < Rank.AllPlayerInfo[j+1].Roomid{
				temp := Rank.AllPlayerInfo[j]
				Rank.AllPlayerInfo[j] = Rank.AllPlayerInfo[j+1]
				Rank.AllPlayerInfo[j+1] = temp
				flag = false
			}
			if flag{
				break
			}
		}
	}
	//封装成AllPlayer类型。
	Sendata := &Data.RoomPlayInfo{
		AllPlayerInfo: Rank.AllPlayerInfo,
		PlayerLen:     NowLen,
	}
	fmt.Println("房间号为",nowRoomId,"更新了积分信息为:",Sendata)
	buf,_ := proto.Marshal(Sendata)
	msg := snet.NewMsgPackage(UPDATA_FEN_ACK,buf)
	NowRoom.BroadRoom(msg)
}
