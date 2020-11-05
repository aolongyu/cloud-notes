package snet

//Create By Cdw 7.21
import (
	"fmt"
	"time"

	"Settings"
	"sync"
)

var RoomMgr RoomMannger

//存放房间的Map和方法
type RoomMannger struct {

	//房间号所对应的房间
	AllRoom map[int32]*Room

	//创建房间需要的锁。不能创建相同ID的房间号
	sync.RWMutex

	//房间

}

func NewRoomMannger() *RoomMannger {
	RoomMgr = RoomMannger{
		AllRoom: make(map[int32]*Room),
	}
	return &RoomMgr
}

func (r *RoomMannger) AddRoom(id int32) int32 {
	r.Lock()
	time.Sleep(200 * time.Millisecond)
	if _, ok := r.AllRoom[id]; !ok {
		fmt.Println("id号为", id, "的房间被创建")
		//Logs.Debug("id号为", id, "的房间被创建")
		r.AllRoom[id] = NewRoom(id)
		r.Unlock()
	} else {
		return id
		r.Unlock()
		id++
		return r.AddRoom(id)
	}
	return id
}

//访问房间ID中的room，存在就返回Room的信息
func (r *RoomMannger) GetRoom(id int32) *Room {
	value, ok := r.AllRoom[id]
	if !ok {
		return nil
	}
	return value
}

//获得部分房间
func (r *RoomMannger) GetAllRoom() map[int32]*Room {
	return r.AllRoom

}
func (r *RoomMannger) EnterRoom(id int32, player PlayerConn) bool {
	r.Lock()
	defer r.Unlock()

	room := r.GetRoom(id)
	if room.GetLen() < Settings.GlobalObject.PlayerNum {
		room.AddPlay(player)
		fmt.Println("[info]编号为", player.GetPlayerId(), "的玩家 ， 进入了房间号为", id, "的房间,当前玩家数为:", room.GetLen())
		//Logs.Debug("编号为 ", player.GetPlayerId(), " 的玩家 ， 进入了房间号为 ", id, " 的房间,当前玩家数为:", room.GetLen())
		//1表示加入房间成功
		player.SendMessage(8, []byte("1"))
		return true
	} else {

		//0表示加入房间失败
		player.SendMessage(8, []byte("0"))
		fmt.Println("玩家", player.GetPlayerId(), "不可加入房间ID号：", id, ",当前玩家数为", room.GetLen())
		//Logs.Debug("玩家",player.GetPlayerId(),"不可加入房间ID号：", id,",当前玩家数为",room.GetLen())
		return false
	}
}
