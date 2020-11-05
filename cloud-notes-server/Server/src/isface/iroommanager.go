package isface

//存放房间的Map和方法
type IRoomManger interface {

	//添加房间，返回房间ID
	AddRoom(id int32) int32

	//根据房间ID获取房间信息
	GetRoom(id int32) IRoom

	//进入房间号为ID，进入玩家
	EnterRoom(id int32, player IPlayer)
	//得到全部房间
	GetAllRoom() map[int32]IRoom
}
