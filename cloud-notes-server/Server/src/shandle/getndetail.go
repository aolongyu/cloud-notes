package shandle

import (
	"isface"
	"snet"
)

type GetNdetail struct{
	snet.BaseRouter
}


func(T GetNdetail) Handle(request isface.IRequest){
	conn := request.GetConnection()
	conn.SendMesg([]byte("1"),[]byte("1"))
}