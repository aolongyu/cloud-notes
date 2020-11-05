package isface

//客户端发来的请求信息

type IRequest interface {

	//获取发送请求的链接
	GetConnection() IConnection

	//获取请求的数据
	GetData() []byte

	//获取请求的消息ID
	GetMsgId() uint32
}
