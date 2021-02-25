package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type Noteclose struct{
	snet.BaseRouter
}


type NotecloseJson struct{
	Id int `json:"id"`
}

func(T Noteclose)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := NotecloseJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle NotecloseJson 传来的信息：",RecvData)

	Line := snet.SDB.Debug().Exec("call note_close(?)",RecvData.Id).RowsAffected

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

