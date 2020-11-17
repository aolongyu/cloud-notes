package isface

import (
	"net"
)

//基础连接模块的接口


type IConnection interface {
	//启动连接，开始工作。
	Start()

	//停止连接，结束当前连接状态
	Stop()

	//获取当前链接的conn
	GetTcpConnection() *net.TCPConn

	//获取当前连接的ID,也算是后来的游戏玩家ID（唯一)
	GetConnID() uint32

	//设置当前连接的ID
	SetConnID(id uint32)

	//当前连接客户端的IP地址
	GetAddr() net.Addr

	//直接把服务器上面的信息发送给TCP客户端
	SendMesg(nums []byte,data []byte) error

	////获取当前roomManager
	//GetRoomManager() IRoomManger
}