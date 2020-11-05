package snet

import (
	"isface"
	"time"
)

const (
	NORMAL = iota
	READY
	PLAY
)

//链接管理

var ConnMap map[uint32]*ConnState = make(map[uint32]*ConnState)

type ConnState struct {
	Conn   isface.IConnection
	Count  int32 //进入对局后的心跳包，平时为-1，进入对局后玩家每发一条消息，+1
	State  int32 //0正常态，1准备态，2对局态
	Roomid int32 //0没有房间。！0有房间
}

func Times() {
	for {
		time.Sleep(6 * time.Second)
		for k, v := range ConnMap {
			if v.Count == 0 {
				v.Conn.Stop()
				delete(ConnMap, k)
			} else if v.Count > 0 {
				v.Count = 0
			}
		}
	}
}

func (co *ConnState) DisConn() {
	//	这是在退出时调用的方法正常退出与非正常退出都应该考虑
	//正常退出---链接状态为0.count为-1，Roomid为0
	//非正常退出---链接状态为1或2，count为[-1,任意],Roomid为0或有值
	if co.State == 0 && co.Roomid == 0 && co.Count == -1 {
		//正常退出
		delete(ConnMap, co.Conn.GetConnID())
	} else {
		//非正常退出
		//先找到该链接的房间
		connroomid := co.Roomid
		if connroomid != 0 && RoomMgr.GetRoom(connroomid) != nil {
			//如果该链接有房间
			nowroom := RoomMgr.GetRoom(connroomid)
			if co.State != 0 {
				//该链接的状态还是非0时，从房间内删除该玩家
				time.Sleep(3 * time.Second)
				delete(nowroom.AllPlayer, int32(co.Conn.GetConnID()))
			}
			if len(nowroom.AllPlayer) == 0 {
				nowroom.ChangeChan <- GAME_STATUS_OVER
				time.Sleep(3 * time.Second)
			}

		}
		delete(ConnMap, co.Conn.GetConnID())
	}
}
func AddConn(uid uint32, conn isface.IConnection) {
	//这是在刚登录时调用的方法
	//添加一个链接到链接模块 Connm[UId]ConnState
	ConnMap[uid] = &ConnState{
		Conn:   conn,
		Count:  -1,
		State:  NORMAL,
		Roomid: 0,
	}
}
func (co *ConnState) UpdateExitRoom() {
	//这是在房间内退出房间调用的方法
	//若是房主退出导致房间销毁，所有人都会修改状态与Roomid
	//在这只考虑修改个人的链接信息，房主退出时，遍历房间内玩家注意调用该方法
	co.State = NORMAL
	co.Roomid = 0
}
func (co *ConnState) Heartbeats() {
	//这是客户端每更新一次数据，心跳加1
	co.Count += 1
}
func (co *ConnState) UpdateJoinRoom(rid int32) {
	//这是链接进入房间后的状态改变
	//State = 1---进入准备态
	co.State = READY
	//Roomid = rid---绑定房间号码
	co.Roomid = rid
}
func (co *ConnState) UpdateGameDoing() {
	//游戏开始时将链接状态的转变
	//进入游戏态
	co.State = PLAY
	//心跳包开始
	co.Count = 1
}
func (co *ConnState) UpdateGameOver() {
	//对局结束后，恢复房间内所有玩家的状态
	//State  = 0--正常态
	co.State = NORMAL
	//Count  = -1---心跳包失效
	co.Count = -1
	//Roomid = 0--取消房间绑定
	co.Roomid = 0
}
