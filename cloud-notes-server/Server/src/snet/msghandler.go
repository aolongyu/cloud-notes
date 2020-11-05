package snet

/*
	负责存放函数与id的键值对
*/
import (
	"Settings"
	"fmt"
	"isface"
	"strconv"
)

type MsgHandle struct {
	//存放Handle的Map
	HandleMap map[uint32]isface.IRouter

	//负责Worker取任务的消息队列
	MessageQueue []chan isface.IRequest

	//业务工作Worker池的数量
	WorkerPoolSize int32

	IsRoomHandleMap map[uint32]bool
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		HandleMap:      make(map[uint32]isface.IRouter),
		WorkerPoolSize: Settings.GlobalObject.WorkerPoolSize,
		MessageQueue:   make([]chan isface.IRequest, Settings.GlobalObject.WorkerPoolSize),
		IsRoomHandleMap:make(map[uint32]bool),
	}
}

//处理消息
func (msgh *MsgHandle) DoMsgHandler(request isface.IRequest) {
	handler, ok := msgh.HandleMap[request.GetMsgId()]
	if !ok {
		fmt.Println("MsgId错误，id= ", request.GetMsgId(), "不存在")
		//Logs.Error("MsgId错误，id= ",request.GetMsgId(),"不存在")

		return
	}
	handler.Handle(request)
}

//为msgHanler 添加处理逻辑的方法
func (msgh *MsgHandle) AddRouter(msgId uint32, router isface.IRouter,detail string,value int32) {
	//判断当前msgid是否绑定了处理方法
	if _, ok := msgh.HandleMap[msgId]; ok {
		panic("不能添加这个id的api了，更换一个 = :" + strconv.Itoa(int(msgId)))
		//Logs.Error("不能添加这个id的api了，更换一个 = :"+ strconv.Itoa(int(msgId)))
	}

	//添加进去
	msgh.HandleMap[msgId] = router
	fmt.Println("添加api成功，id为", msgId,"添加内容为：",detail)
	//Logs.Debug("添加api成功，id为",msgId,"添加内容为:",detail)
	if value == 0{
		//说明不是房间Handle
		msgh.IsRoomHandleMap[msgId] = false
	}else{
		//说明是房间Handle
		msgh.IsRoomHandleMap[msgId] = true
	}
}

//启动一个Worker工作池（开启工作池的动作只能发生一次，一个框架只开一次）
func (msgh *MsgHandle) StartWorkerPool() {
	//根据workerPoolSize 分别开启Worker,每个Worker用一个go去承载
	for i := 0; i < int(msgh.WorkerPoolSize); i++ {
		//一个worker被启动
		//1 当前worker对应的channel消息队列，开启空间
		msgh.MessageQueue[i] = make(chan isface.IRequest, Settings.GlobalObject.MaxWorkerSize)

		//启动当前worker，阻塞等待消息从channer传递进来
		go msgh.StartOneWorker(int32(i), msgh.MessageQueue[i])
	}
}

//启动一个Worker工作流程
func (msgh *MsgHandle) StartOneWorker(workerId int32, taskQueue chan isface.IRequest) {
	fmt.Println("[info]工作池", workerId, "开始工作 ...")
	//Logs.Debug("[info]工作池",workerId,"开始工作 ...")
	//不断阻塞等待对应的消息队列
	for {
		select {
		//如果有消息过来，出列就是客户端的Request，执行当前Request绑定的业务
		case request := <-taskQueue:
			msgh.DoMsgHandler(request)
		}
	}
}

func(msgh *MsgHandle) SendMsgToMesQueue(request isface.IRequest) {

	//均分给所有worker
	msgQue := msgh.ReturnChannel(request)
	msgQue <- request

}

func(msgh *MsgHandle)ReturnChannel(request isface.IRequest) chan isface.IRequest{
	//说明是房间方法，那么返回房间号的管道
	ConnId := request.GetConnection().GetConnID()
	if msgh.IsRoomHandleMap[request.GetMsgId()] {
		//TODO 返回房间管道
		romid := ConnMap[ConnId].Roomid
		if _,ok := RoomMgr.AllRoom[romid];!ok{
			fmt.Println(ConnMap[ConnId].Roomid,"不存在")
			workid := int32(request.GetConnection().GetConnID()  ) % msgh.WorkerPoolSize
			return msgh.MessageQueue[workid]
		}
		NowRoom := RoomMgr.GetRoom(int32(ConnMap[ConnId].Roomid))
		return NowRoom.TaskQueue
	}else{//说明不是房间方法，那么返回工作池id即可
		workid := int32(ConnId ) % msgh.WorkerPoolSize
		return msgh.MessageQueue[workid]
	}
	//如果忘记注册的话，还是返回工作池队列
	workid := int32(request.GetConnection().GetConnID()  ) % msgh.WorkerPoolSize
	return msgh.MessageQueue[workid]
}