package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CollectNoteBookAdd struct{
	snet.BaseRouter
}
//收藏笔记本ID
type CollectNoteBookaddJson struct{
	noteid int `json:"noteid"`
	bid   int `json:"bid"`
}


func(T CollectNoteBookAdd)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := CollectNoteBookaddJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle CollectNoteBookAdd 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call collectNoteBook_add(?,?)",RecvData.noteid,RecvData.bid).RowsAffected

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
