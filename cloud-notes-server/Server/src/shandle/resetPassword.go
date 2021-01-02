package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ResetPassword struct{
	snet.BaseRouter
}

type ResetPasswordJson struct{
	Id int `json:"id"`

}

func(T ResetPassword)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := ResetPasswordJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle ResetPassword 传来的信息：",RecvData)

	Line := snet.SDB.Debug().Exec("call resetPassword(?)",RecvData.Id).RowsAffected

	fmt.Println("Line : ",Line)
	res := Status{}
	if Line > 0{
		res.Status = "1"
		SendData,_ := json.Marshal(res)
		conn.SendMesg([]byte(""),SendData)
	}else{
		res.Status = "0"
		SendData,_ := json.Marshal(res)
		conn.SendMesg([]byte(""),SendData)
	}
}
