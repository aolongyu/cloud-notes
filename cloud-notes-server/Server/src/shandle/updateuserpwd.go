package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type UpdateUserImf struct{
	snet.BaseRouter
}

type UpdateUserImfJson struct{
	Uid int `json:"uid"`
	Pwd string `json:"pwd"`
}

func(T UpdateUserImf)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := UpdateUserImfJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle UpdateUserImfJson 传来的信息：",RecvData)

	Line := snet.SDB.Debug().Exec("call update_pwd(?,?)",RecvData.Uid,RecvData.Pwd).RowsAffected

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


