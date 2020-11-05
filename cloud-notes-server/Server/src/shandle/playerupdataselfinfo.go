package shandle

import (
	Data "data"
	//"fmt"
	"isface"
	"snet"
	//"time"

	"github.com/golang/protobuf/proto"
)

type PlayerUpdataSelfInfo struct {
	snet.BaseRouter
}

func (p *PlayerUpdataSelfInfo) Handle(request isface.IRequest) {

	PlayerMsgInfo := &Data.PlayerPro{
		PlayerMsg:   nil,
		PlayerCoord: nil,
		PlayerHero:  nil,
	}
	proto.Unmarshal(request.GetData(), PlayerMsgInfo)

	NowRoom := snet.RoomMgr.GetRoom(PlayerMsgInfo.PlayerMsg.Roomid)
	NowRoom.Lock()
	defer NowRoom.Unlock()
	UpdataPlayer := snet.PlayerConn{
		Player: snet.Player{
			Coord: snet.Coord{
				X: PlayerMsgInfo.PlayerCoord.X,
				Y: PlayerMsgInfo.PlayerCoord.Y,
			},
			Id: PlayerMsgInfo.PlayerMsg.Uid,
			BindHero: snet.BindHero{
				HeroID:      PlayerMsgInfo.PlayerHero.HeroID,
				Skill:       PlayerMsgInfo.PlayerHero.Skill,
				Weapon:      PlayerMsgInfo.PlayerHero.Weapon,
				HeroLv:      PlayerMsgInfo.PlayerHero.HeroLv,
				HeroHp:      PlayerMsgInfo.PlayerHero.HeroHp,
				AttBehavior: PlayerMsgInfo.PlayerHero.AttBehavior,
				RunBehavior: PlayerMsgInfo.PlayerHero.RunBehavior,
				HeroAtk:     PlayerMsgInfo.PlayerHero.HeroAtk,
				HeroSpe:     PlayerMsgInfo.PlayerHero.HeroSpe,
				HeroDef:     PlayerMsgInfo.PlayerHero.HeroDef,
				HeroVector:  PlayerMsgInfo.PlayerHero.HeroVector,
			},
			UserName: PlayerMsgInfo.PlayerMsg.Username,
			RoomId:   PlayerMsgInfo.PlayerMsg.Roomid,
		},
		IConnection: request.GetConnection(),
	}
	//fmt.Println("时间戳为:", time.Now().UnixNano()/1e6, "玩家ID为：", PlayerMsgInfo.PlayerMsg.Uid, "的玩家更新了自己的信息", PlayerMsgInfo)
	snet.ConnMap[uint32(PlayerMsgInfo.PlayerMsg.Uid)].Heartbeats()
	//给更新的+
	NowRoom.UpdataPlayerMap[PlayerMsgInfo.PlayerMsg.Uid] = UpdataPlayer
	//给全局的＋
	NowRoom.AllPlayer[PlayerMsgInfo.PlayerMsg.Uid] = UpdataPlayer
}
