package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type DeleteUserById struct {
	snet.BaseRouter
}

func(T DeleteUserById)Handle(request isface.IRequest){
	conn := request.GetConnection()
	recvData := ViewNoteJson{}

	json.Unmarshal(request.GetData(),&recvData)

	fmt.Println("Handle DeleteByUserId 接收到消息",recvData)

	Data := Status{}

	Line := snet.SDB.Debug().Raw("call delete_user(?)",recvData.Id).RowsAffected

	if Line > 0{
		Data.Status = "1"
	}else{
		Data.Status = "0"
	}
	SendData,_ := json.Marshal(Data)
	conn.SendMesg([]byte(""),SendData)
}