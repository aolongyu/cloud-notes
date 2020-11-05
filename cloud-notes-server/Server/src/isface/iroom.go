package isface

type IRoom interface {
	//设置房间值
	SetValue(value int32)

	//改变房间状态
	ChangeState()

	//广播
	BroadRoom(message IMessage)

	//获取int和play的键值对
	GetAllPlay() map[int32]IPlayer

	//获得该房间的长度
	GetLen() int32

	//获取房间id
	GetRoomID() int32

	//获取房间状态
	GetRoomStatus() int32

	//插入一个玩家
	AddPlay(player IPlayer)
}

//房间状态类
type RoomState interface {
	Handle(request IRequest)
}
