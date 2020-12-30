package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type EditNote struct{
	snet.BaseRouter
}

type EditNoteJson struct{
	Note_id int `json:"Note_id"`
	Note_name string `json:"Note_name"`
	Note_introduction string `json:"introduction"`
	Note_type int `json:"Note_type"`
	Note_text string `json:"Note_text"`
}

func(T EditNote)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := EditNoteJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle EditNote 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call editnote(?,?,?,?,?)",RecvData.Note_id,RecvData.Note_name,RecvData.Note_introduction,RecvData.Note_type,RecvData.Note_text).RowsAffected

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
