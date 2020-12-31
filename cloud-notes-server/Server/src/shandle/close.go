package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CloseUserByid struct {
	snet.BaseRouter
}
//用户id
type CloseUserByidJson struct{
	Uid int `json:"Uid"`
}


func(T CloseUserByid)Handle(request isface.IRequest){
	conn:=request.GetConnection()
	recvData := CloseUserByidJson{}

	json.Unmarshal(request.GetData(),&recvData)

	fmt.Println("Handle CloseUserByName 接收到信息：",recvData)

	Line := snet.SDB.Debug().Raw("call close(?)",recvData.Uid).RowsAffected

	Data := Status{}

	if Line > 0{
		Data.Status = "1"
	}else{
		Data.Status = "0"
	}
	SendData,_ := json.Marshal(Data)
	conn.SendMesg([]byte(""),SendData)
}