package snet

import (
	"isface"
	)

type Request struct {

	//与客户端建立好的连接
	conn isface.IConnection

	//客户端发来的数据
	msg isface.IMessage
}

//获取请求连接信息

func (rs *Request)GetConnection() isface.IConnection{
	return rs.conn
}

//获取请求的数据
func (rs *Request) GetData() []byte{
	return rs.msg.GetData()
}

//获取请求的消息ID
func (rs *Request) GetMsgId()uint32 {
	return rs.msg.GetMsgId()
}
