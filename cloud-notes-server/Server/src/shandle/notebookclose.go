package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type NoteBookclose struct{
	snet.BaseRouter
}

type NoteBookcloseJson struct{
	Id int `json:"id"`
}

func(T NoteBookclose)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := NoteBookcloseJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle NoteBookcloseJson 传来的信息：",RecvData)

	Line := snet.SDB.Debug().Exec("call notebook_close(?)",RecvData.Id).RowsAffected

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

