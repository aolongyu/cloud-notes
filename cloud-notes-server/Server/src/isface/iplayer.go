package isface


type IPlayer interface {
	//获得玩家ID
	GetPlayerId()int32

	//获得玩家链接
	//GetConnId()IConnection

	//发送信息
	SendMessage(id uint32 , data []byte)

	GetPlayerName() string

}
