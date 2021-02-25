package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type DelCBook struct{
	snet.BaseRouter
}

type DelCBookJson struct{
	bid   int `json:"bid"`
}


func(T DelCBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := DelCBookJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle DelCBook 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call delete_cbook(?)",RecvData.bid).RowsAffected

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
