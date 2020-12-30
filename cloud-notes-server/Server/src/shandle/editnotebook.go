package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type EditNoteBook struct{
	snet.BaseRouter
}

type EditNoteBookJson struct{
	Note_id int `json:"Id"`
	Note_name string `json:"notebook_name"`
	Note_introduction string `json:"introduction"`
	Note_type int `json:"notebook_type"`
}

func(T EditNoteBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := EditNoteBookJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle EditNoteBook 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call Editnotebook(?,?,?,?)",RecvData.Note_id,RecvData.Note_name,RecvData.Note_introduction,RecvData.Note_type).RowsAffected

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
