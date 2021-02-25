package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CollectNoteBookRemove struct{
	snet.BaseRouter
}

type CollectNoteBookRemoveJson struct{
	noteid int `json:"noteid"`
}


func(T CollectNoteBookRemove)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := CollectNoteBookRemoveJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle CollectNoteBookRemoveJson 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call collectNoteBook_remove(?)",RecvData.noteid).RowsAffected

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
