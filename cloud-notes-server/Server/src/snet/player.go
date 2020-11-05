//Create By Cdw 7.21
package snet

import (
	Data "data"
	"isface"
)

//坐标类
type Coord struct{
	X float64
	Y float64
}

//玩家链接信息
type PlayerConn struct{
	Player
	isface.IConnection
}

//绑定英雄的属性
type BindHero struct {
	HeroID int32
	Skill int32
	Weapon int32
	HeroLv int32
	HeroHp int32
	//攻击行为
	AttBehavior int32
	//移动行为
	RunBehavior int32
	//没啥用
	HeroAtk int32
	HeroSpe int32
	HeroDef int32
	//方向
	HeroVector float64
}

//玩家基本信息
type Player struct{
	Coord
	Id int32 // 玩家ID虽然和Connid一样了
	BindHero
	UserName string
	RoomId int32
}

func (p PlayerConn)GetPlayerId()int32{
	return p.Id
}
//
//func (p *PlayerConn) GetConnId() isface.IConnection{
//	return p.GetConnID()
//}

func (p PlayerConn) SendMessage(id uint32,data []byte){
	p.SendMesg(id,data)
}

func (p PlayerConn)GetPlayerName() string {
	return p.UserName
}

func NewPlayerConn(PlayerId int32,roomId int32,connection isface.IConnection) *PlayerConn {
	a := &PlayerConn{
		Player:      Player{
			Coord:  Coord{
				X: 0,
				Y: 0,
			},
			Id:     PlayerId,
			UserName:"",
			RoomId: roomId,
		},
		IConnection: connection,
	}
	return a
}

func (p *PlayerConn)SetCoord(X float64,Y float64){
	Temp := RoomMgr.GetRoom(p.RoomId).AllPlayer[p.Id]
	Temp.X = X
	Temp.Y = Y
	RoomMgr.GetRoom(p.RoomId).AllPlayer[p.Id] = Temp
}
func (p *PlayerConn)SetPlayerMsg(heroID int32,skill int32,heroLV int32,heroWeapon int32){
	Temp := RoomMgr.GetRoom(p.RoomId).AllPlayer[p.Id]
	Temp.Player.HeroID = heroID
	Temp.Player.Skill = skill
	Temp.Player.HeroLv = heroLV
	Temp.Player.Weapon = heroWeapon
	RoomMgr.GetRoom(p.RoomId).AllPlayer[p.Id] = Temp
}

func (p *PlayerConn)GetPlayAllMsg() *Data.PlayerPro{
	value := &Data.PlayerPro{
		PlayerMsg:   &Data.PlayerInfo{
			Roomid:   p.RoomId,
			Uid:      p.Id,
			Username: p.UserName,
		},
		PlayerCoord: &Data.CoordPro{
			X: p.X,
			Y: p.Y,
		},
		PlayerHero:  &Data.HeroPro{
			Skill:       p.Skill,
			Weapon:      p.Weapon,
			HeroLv:      p.HeroLv,
			HeroHp:      p.HeroHp,
			HeroAtk:     p.HeroAtk,
			HeroSpe:     p.HeroSpe,
			HeroDef:     p.HeroDef,
			HeroVector:  p.HeroVector,
			AttBehavior: p.AttBehavior,
			RunBehavior: p.RunBehavior,
			HeroID:      p.HeroID,
		},
	}
	return value
}

func (p *PlayerConn)SetPlayerUserName(username string){
	Temp := RoomMgr.GetRoom(p.RoomId).AllPlayer[p.Id]
	Temp.UserName = username
	RoomMgr.GetRoom(p.RoomId).AllPlayer[p.Id] = Temp
}