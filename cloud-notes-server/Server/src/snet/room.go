package snet

//Create By Cdw 7.21
import (
	"Settings"
	Data "data"
	"encoding/binary"
	"fmt"

	"isface"
	"reflect"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	GAME_STATUS_READY = iota
	GAME_STATUS_BATTLE
	GAME_STATUS_OVER
)

//存放房间的Map方法
type Room struct {

	//房间ID
	Id int32

	//房间状态
	State RoomState

	//游戏状态值
	GameValue int32

	//玩家id与玩家对应，游戏开始后房主预处理所有玩家位置
	AllPlayer map[int32]PlayerConn

	//上锁否则加入的时候会出问题
	sync.RWMutex

	//全部地图资源的集合游戏开始后，游戏开始后房主预处理所有物品
	AllResource map[Coord]MapResource

	//用来更新玩家信息的。
	UpdataPlayerMap map[int32]PlayerConn
	//各种管道控制定时的。方法较差如果改进可看时间轮算法。
	//用于让GameReady   和 GameBattle   和 GAME_OVER时候定时执行一些方法。
	ChangeChan chan int32 //用于改变状态的管道

	DoHandleChan chan bool //用于执行DoHandle

	CanDoHandle bool //用来判断是否可以用定时器进行发送

	//房间任务池
	TaskQueue chan isface.IRequest

	msgh isface.IMsgHandle

	Ranking *Data.RoomPlayInfo
}

func NewRoom(value int32) *Room {
	a := &Room{
		Id:              value,
		State:           &GameReady{},
		GameValue:       0,
		AllPlayer:       make(map[int32]PlayerConn),
		AllResource:     make(map[Coord]MapResource),
		ChangeChan:      make(chan int32, 2),
		DoHandleChan:    make(chan bool, 1),
		UpdataPlayerMap: make(map[int32]PlayerConn),
		CanDoHandle:     true,
		TaskQueue:       make(chan isface.IRequest, Settings.GlobalObject.WorkerPoolSize),
		msgh:            GlobalServer.msgHandler,
		Ranking:         new(Data.RoomPlayInfo),
	}
	a.ChangeState()
	///一直执行定时器内容
	time.Sleep(1 * time.Second)
	go a.DoStateHandle()
	go a.DoRoomHandle()
	return a
}

//房间专门开一个go去处理队列中的信息
func (r *Room) DoRoomHandle() {
	fmt.Println("房间号为：", r.GetRoomID(), "的工作池开启")
	for {
		select {
		case re := <-r.TaskQueue:
			r.msgh.DoMsgHandler(re)
		}
		if r.GameValue == GAME_STATUS_OVER {
			return
		}
	}
}

func (r *Room) SetValue(value int32) {
	r.GameValue = value
	r.ChangeState()
}

func (r *Room) ChangeState() {

	if r.GameValue == GAME_STATUS_READY { //游戏准备状态

		r.State = &GameReady{}
		fmt.Println("[info]房间号为", r.GetRoomID(), "游戏进入准备状态", reflect.TypeOf(r.State))
		//Logs.Debug("[info]房间号为", r.GetRoomID(), "游戏进入准备状态", reflect.TypeOf(r.State))
	} else if r.GameValue == GAME_STATUS_BATTLE { //游戏正在进行状态

		r.State = &GameDoing{}
		fmt.Println("[info]房间号为", r.GetRoomID(), "游戏进入进行状态", reflect.TypeOf(r.State))
		//Logs.Debug("[info]房间号为", r.GetRoomID(), "游戏进入进行状态", reflect.TypeOf(r.State))
	} else if r.GameValue == GAME_STATUS_OVER { //游戏结束状态

		r.State = &GameEnd{}
		fmt.Println("[info]房间号为", r.GetRoomID(), "游戏结束", reflect.TypeOf(r.State))
		data := &Data.PlayerInfo{
			Roomid:   r.GetRoomID(),
			Uid:      0,
			Username: "",
		}
		buf, _ := proto.Marshal(data)
		data1 := &Request{
			conn: nil,
			msg:  NewMsgPackage(1, buf),
		}
		//当房间结束后，删除房间。
		r.State.Handle(data1)
		//Logs.Debug("[info]房间号为", r.GetRoomID(), "游戏结束", reflect.TypeOf(r.State))
	}

}

//对该房间广播
func (r Room) BroadRoom(message isface.IMessage) {
	for _, value := range r.AllPlayer {
		//将消息转发给所有人 id = 40 ,把message给所有人发送包括自己
		value.SendMessage(message.GetMsgId(), message.GetData())
	}
}

//获取int和play的键值对
func (r Room) GetAllPlay() map[int32]PlayerConn {
	return r.AllPlayer
}

//获取房间长度
func (r Room) GetLen() int32 {
	return int32(len(r.AllPlayer))
}

//获取房间id
func (r Room) GetRoomID() int32 {
	return r.Id
}

//获取房间状态
func (r Room) GetRoomStatus() int32 {
	return r.GameValue
}
func (r Room) AddPlay(player PlayerConn) {
	r.AllPlayer[player.GetPlayerId()] = player
}

var Hearts int32 = 60 //60个100ms
func (r Room) Tick(timeNeed int32) {
	for {
		time.Sleep(time.Duration(timeNeed) * time.Millisecond)
		if r.GameValue == GAME_STATUS_BATTLE && Hearts == 0 {
			for k, _ := range r.AllPlayer {
				if ConnMap[uint32(k)].Count == 0 {
					ConnMap[uint32(k)].Conn.Stop()
					delete(ConnMap, uint32(k))
				} else if ConnMap[uint32(k)].Count > 0 {
					ConnMap[uint32(k)].Count = 0
				}
			}
			Hearts = 60
		}
		Hearts--
		r.DoHandleChan <- true
		if r.GameValue == int32(GAME_STATUS_OVER) {
			return
		}
	}
}

//不断执行这个handle
func (r *Room) DoStateHandle() {
	var timeslice int32 = Settings.GlobalObject.Time
	go r.Tick(timeslice)
	//房间最开始的时间
	var counts int32 = 444444 / timeslice
	var GameTime int32 = 40003

	PlayerInfoData := &Data.PlayerInfo{
		Roomid:   r.GetRoomID(),
		Uid:      0,
		Username: "",
	}
	buf,_ := proto.Marshal(PlayerInfoData)
	RequestMsg := &Request{
		conn: nil,
		msg:  NewMsgPackage(1, buf),
	}
	for {
		select {
		case s := <-r.ChangeChan:

			time.Sleep(10 * time.Millisecond)
			r.SetValue(s)
			time.Sleep(30 * time.Millisecond)
			//如果s等于Game_status_over那么就退出这个阻塞的协程
			if s == GAME_STATUS_OVER {
				return
			}
			if s == GAME_STATUS_BATTLE {
				//设置游戏进行的时间
				counts = 1000 / timeslice
				GameTime = 180
				fmt.Println(r.GetRoomID(), "的房间将开始游戏，持续", GameTime)
			}
		case <-r.DoHandleChan:
			counts--
			//fmt.Println(  reflect.TypeOf(r.State))
			if r.CanDoHandle {
				RequestMsg.msg.SetMsgId(1)
				r.State.Handle(RequestMsg)
			}
			//时间包
			if counts == 0 && GameTime != 0 {
				counts = 1000 / timeslice
				GameTime--
				RequestMsg.msg.SetMsgId(2)
				r.State.Handle(RequestMsg)
			}
			//说明一局已经结束了，那么切换游戏状态
			if GameTime == 0 {
				fmt.Println("游戏结束，准备关闭房间")
				GameTime--
				//Logs.Debug("房间",r.GetRoomID(),"游戏结束")
				r.CanDoHandle = false
				RequestMsg.msg.SetMsgId(3)
				r.State.Handle(RequestMsg)
			}
		}
	}
	return
}

//房间状态接口
type RoomState interface {
	Handle(request isface.IRequest)
}

//房间准备状态类
type GameReady struct {
}

func (g *GameReady) Handle(request isface.IRequest) {
	//fmt.Println("Ready!!")

}

//游戏进行中的类
type GameDoing struct {
}

//游戏进行中的Handle,将内容敌人信息广播
func (g *GameDoing) Handle(request isface.IRequest) {
	//fmt.Println("Doing!!")
	rom := &Data.PlayerInfo{}
	proto.Unmarshal(request.GetData(), rom)
	if request.GetMsgId() == uint32(1) {
		UpdataAllPlayerResorce(rom.Roomid)
	} else if request.GetMsgId() == uint32(2) {
		CountDown(rom.Roomid)
	} else if request.GetMsgId() == uint32(3) {
		RoomMgr.GetRoom(rom.Roomid).ChangeChan <- GAME_STATUS_OVER
	}
}

//游戏结束时候的类
type GameEnd struct {
}

//游戏结束时候执行的Handle
func (g *GameEnd) Handle(request isface.IRequest) {
	fmt.Println("Over!!")
	msg := &Data.PlayerInfo{
		Roomid:   0,
		Uid:      0,
		Username: "",
	}
	proto.Unmarshal(request.GetData(), msg)
	//回收资源，关闭所有协程
	if request.GetMsgId() == uint32(1) {
		CloseRoom(msg.Roomid)
	}
}
func CloseRoom(rid int32) {
	//把房间内的所有人都踢出来，关房间，设置链接心跳包初始状态
	nowroom := RoomMgr.AllRoom[rid]
	//处理房间积分
	DealRank(rid)
	fmt.Println("处理房间积分，房间号为：",rid)
	for k, _ := range nowroom.AllPlayer {

		ConnMap[uint32(k)].UpdateGameOver()
	}
	delete(RoomMgr.AllRoom, rid)
}
func UpdataAllPlayerResorce(roomid int32) {

	NowRoom := RoomMgr.AllRoom[roomid]
	//var aa []Data.MapResourcePro
	//var i int = 0
	var ExitChan chan bool = make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(15) * time.Millisecond)
		ExitChan <- true
	}()
	for {
		if len(NowRoom.UpdataPlayerMap) == len(NowRoom.AllPlayer) {
			NowRoom.UpdataPlayerMap = make(map[int32]PlayerConn,len(NowRoom.AllPlayer))
			NowRoom.CanDoHandle = false
			goto LOOP
		}
		select {
		case <-ExitChan:
			goto lable
		}
	}
lable:
	NowRoom.UpdataPlayerMap = make(map[int32]PlayerConn,len(NowRoom.AllPlayer))
	NowRoom.CanDoHandle = false
LOOP:
	NowRoom.RLock()
	defer NowRoom.RUnlock()
	LenR := len(NowRoom.AllResource)
	resources := make([]*Data.MapResourcePro, LenR)
	i := 0
	for k, v := range NowRoom.AllResource {
		resources[i] = &Data.MapResourcePro{
			Roomid: roomid,
			Id:     v.Id,
			MapResourceCoord: &Data.CoordPro{
				X: k.X,
				Y: k.Y,
			},
		}
		i++
	}

	//var bb []Data.PlayerPro
	LenP := len(NowRoom.AllPlayer)
	palys := make([]*Data.PlayerPro, LenP)
	j := 0
	for _, vv := range NowRoom.AllPlayer {
		palys[j] = vv.GetPlayAllMsg()
		j++
	}
	pre := &Data.PreloadPro{
		AllPlayer:      palys,
		AllMapResource: resources,
	}
	time.Sleep(time.Duration(25) * time.Millisecond)
	//fmt.Println("时间戳为：", time.Now().UnixNano()/1e6, "给房间号为", roomid, "的房间广播了信息", pre)
	//Logs.Debug("时间戳为：",time.Now().UnixNano()/1e6,"给房间号为", roomid, "的房间广播了信息", pre)
	msgPre, _ := proto.Marshal(pre)
	//const JOIN_ROOM_BORAD = 10 //给所有人广播加入的人的信息
	msg := NewMsgPackage(101, msgPre)

	NowRoom.BroadRoom(msg)
	NowRoom.CanDoHandle = true
}

func CountDown(roomid int32) {

	NowRoom := RoomMgr.AllRoom[roomid]
	var buf = make([]byte, 4)
	//大端传输可能会有问题C# 默认小端。
	binary.BigEndian.PutUint32(buf, uint32(1))
	msg := NewMsgPackage(104, buf)
	NowRoom.BroadRoom(msg)

}

func GameOver(roomid int32) {
	NowRoom := RoomMgr.GetRoom(roomid)
	NowRoom.SetValue(GAME_STATUS_OVER)
	msg := NewMsgPackage(105, []byte("Game_OVER"))
	NowRoom.BroadRoom(msg)
}


//处理游戏结束时候的积分
func DealRank(roomid int32){
	NeedExp := make([]int32,36)
	NeedExp[1] = 0
	NeedExp[2] = 6
	NeedExp[3] = 20
	NeedExp[4] = 41
	NeedExp[5] = 69
	NeedExp[6] = 104
	NeedExp[7] = 146
	NeedExp[8] = 195
	NeedExp[9] = 251
	NeedExp[10] = 314
	NeedExp[11] = 384
	NeedExp[12] = 461
	NeedExp[13] = 545
	NeedExp[14] = 636
	NeedExp[15] = 734
	NeedExp[16] = 839
	NeedExp[17] = 951
	NeedExp[18] = 1070
	NeedExp[19] = 1196
	NeedExp[20] = 1336
	NeedExp[21] = 1491
	NeedExp[22] = 1662
	NeedExp[23] = 1850
	NeedExp[24] = 2057
	NeedExp[25] = 2284
	NeedExp[26] = 2531
	NeedExp[27] = 2801
	NeedExp[28] = 3094
	NeedExp[29] = 3412
	NeedExp[30] = 3755
	NeedExp[31] = 4126
	NeedExp[32] = 4525
	NeedExp[33] = 4953
	NeedExp[34] = 5412
	NeedExp[35] = 5903
	NowRoom := RoomMgr.GetRoom(roomid)

	Rank := NowRoom.Ranking.AllPlayerInfo

	RankLen := len(Rank)
	cus := Data.Account_info{}
	for i:= 0;i<RankLen;i++{
		//提取用户的名称
		cus.UID = Rank[i].Uid
		SDB.Table("account_info").Where("UID = ?",cus.UID).Scan(&cus)
		exp := cus.Account_exp - int32(2*i)+7

		SDB.Debug().Exec("update account_info set account_exp = ? where Uid = ?",exp,cus.UID)
		var Level int32
		for j := 1 ;j<=35;j++{
			if exp <= NeedExp[j] && exp > NeedExp[j-1]{
				Level = int32(j - 1)
				break;
			}
		}
		SDB.Debug().Exec("update account_info set account_lv = ? where Uid = ?",Level,cus.UID)
		fmt.Println("用户",cus.UID,"的积分增值为:",exp,"等级为：",Level)
	}
}