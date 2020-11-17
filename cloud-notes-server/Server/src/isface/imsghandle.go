package isface

//消息管理

type IMsgHandle interface {

	//执行请求
	DoMsgHandler(request IRequest) //执行方法

	//为ID为msgID的编号，添加方法到router中，方法需要重写Router方法
	AddRouter(msgId string,router IRouter,detail string,value int32)

	//启动Worker工作池
	StartWorkerPool()

	//发送信息给工作池
	SendMsgToMesQueue(request IRequest)
}
