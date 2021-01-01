package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CreateCBook struct{
	snet.BaseRouter
}

type CreateCBookJson struct{
	Uid int `json:"uid"`
	Cname string `json:"cname"`
	Cintroduction string `json:"cintroduction"`
	Ctype int  `json:"ctype"`
}

func(T CreateCBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := CreateCBookJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle CreateCBook 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call create_cbook(?,?,?,?)",RecvData.Uid,RecvData.Cname,RecvData.Cintroduction,RecvData.Ctype).RowsAffected

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
