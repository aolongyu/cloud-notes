package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CloseUserByName struct {
	snet.BaseRouter
}
//用户id
type CloseUserByNameJson struct{
	Uid int `json:"Uid"`
}


func(T CloseUserByName)Handle(request isface.IRequest){
	conn:=request.GetConnection()
	recvData := CloseUserByNameJson{}

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