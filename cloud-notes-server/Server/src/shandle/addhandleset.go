package shandle

import (
	"isface"
)

//在这里添加Handle集合
const (
	LOGIN_HANDLE = 1
	TEST_HANDLE = 404
)

func AddHandleInit(s isface.IServer) {
	//s.AddHandle(TEST_HANDLE,&Nofound{},"测试能否连接",0)
	s.AddHandle("login",&Login{},"登陆",0)
	s.AddHandle("regist",&Register{},"注册",0)
}
