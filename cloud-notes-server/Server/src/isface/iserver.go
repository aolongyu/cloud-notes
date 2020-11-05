package isface

//定义服务器的方法

type IServer interface {

	//启动服务器
	Start()

	//停止服务器
	Stop()

	//开启服务器方法
	Serve()

	//注册业务Handle , 消息ID，router ， 细节， 是否交给房间处理
	AddHandle(msgId uint32,router IRouter,detail string,value int32)
}