package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CollectNoteBook struct {
	snet.BaseRouter
}
//收藏笔记本ID
type CollectNoteBookJson struct{
	collectNoteBookId int `json:"collectNoteBookId"`
	UserId 			  int `json:"userid"`
}

func(T CollectNoteBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	AddMessage := CollectNoteBookJson{}

	json.Unmarshal(request.GetData(),&AddMessage)

	fmt.Println("Handle CollectNoteBook 传来的信息：",AddMessage)

	Line := snet.SDBNote.Debug().Exec("call collectNoteBook_add(?,?)",AddMessage.collectNoteBookId,AddMessage.UserId).RowsAffected

	returnres := Status{}
	if(Line > 0){
		returnres.Status = "1"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""),data)
	}else{
		returnres.Status = "0"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""),data)
	}
}