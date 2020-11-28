package snet

import (
	"Settings"
	"github.com/jinzhu/gorm"
	"logA"

	"fmt"
	"isface"
	"net"
)

//var log *logger.
var Logs *logA.FileLogger = logA.NewFileLogger("debug", "debug.log", "../log")

var GlobalServer *Server

//Server服务器实现
type Server struct {
	//用什么IP协议
	IPVersion string

	//服务绑定的IP地址
	IPaddr string

	//服务器绑定的端口号
	Port int

	//当前服务器的消息管理模块，管理所有的键值对
	msgHandler isface.IMsgHandle
}

//创建服务器句柄
func NewServer() isface.IServer {

	Settings.GlobalObject.Reload()

	s := &Server{
		IPVersion:  "tcp4",
		IPaddr:     Settings.GlobalObject.Host,
		Port:       Settings.GlobalObject.TcpPort,
		msgHandler: NewMsgHandle(),
	}
	GlobalServer = s
	return s
}

//启动服务器
func (s *Server) Start() {
	fmt.Println("[START] Server启动,  服务器ip:", s.IPaddr, "  端口号为：", s.Port, "\n版本号:",
		Settings.GlobalObject.Version, "  最大连接数:", Settings.GlobalObject.MaxConn, "  最大包长:", Settings.GlobalObject.MaxPacketSize)
	//Logs.Debug("[START] Server启动,  服务器ip:", s.IPaddr, "  端口号为：", s.Port, "\n版本号:",
	//	Settings.GlobalObject.Version, "  最大连接数:", Settings.GlobalObject.MaxConn, "  最大包长:", Settings.GlobalObject.MaxPacketSize)

	//go Times()
	//go func() {
	//开启消息队列以及工作池
	s.msgHandler.StartWorkerPool()

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPaddr, s.Port))
	if err != nil {
		fmt.Println("建立链接失败")
		//Logs.Error("建立链接失败!")
		return
	}

	//监听服务器地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("监听类型:", s.IPVersion, "  错误:", err)
		//Logs.Error("监听端口失败!")
		return
	}

	//已经监听成功
	fmt.Println("开启了Server 现在正在监听信息")
	//Logs.Debug("服务器开始监听")

	//TODO 下面写一个自动生成唯一ID的方法，目前采用从0自增的方法，可以考虑雪花算法，之后服务器并行接受时候可以使用，但是这个只开一个协程所以不用

	var cid uint32
	cid = 0
	//3 启动server网络连接业务
	for {
		//阻塞等待建立连接请求
		conn, err := listenner.AcceptTCP()

		if err != nil {
			fmt.Println("接受连接发生错误: ", err)
			//Logs.Error("AccentTCP链接出错")
		}

		//ToDo 如果超过最大连接限制，那么就关闭这个新链接
		dealConn := NewConntion(conn, cid, s.msgHandler)

		cid++
		//开个协程处理该链接
		go dealConn.Start()
	}
	//}()
}

//停止服务器
func (s *Server) Stop() {
	fmt.Println("服务器停止运行 ,监听的端口为", s.Port)
	//Logs.Debug("服务器关闭")
}

//开启服务器方法
func (s *Server) Serve() {
	////TODO要打开服务器的地方
	var err error
	defer SDB.Close()
	SDB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ybj_userdb?charset=utf8")
	SDB.SingularTable(true)
	if err != nil {
		fmt.Println("打开数据库失败")
		println(err)
		//Logs.Error("数据库来链接失败")
	}
	fmt.Println("打开数据库成功")
	//Logs.Debug("数据库链接成功")
	s.Start()
	//ToDo 可以在这里添加一些其他方法，后期添加

	//阻塞住
	select {}
}

//注册业务Handle
func (s *Server) AddHandle(msgId string, router isface.IRouter, detail string, value int32) {
	s.msgHandler.AddRouter(msgId, router, detail, value)
}
