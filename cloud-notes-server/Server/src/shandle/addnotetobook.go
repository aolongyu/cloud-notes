package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)
//
type AddNoteToNoteBook struct{
	snet.BaseRouter
}
//把笔记加入笔记本
type AddJson struct{
	//把笔记加入笔记本，笔记id，笔记本id
	nid int `json:"nid"`
	bid int `json:"bid"`
}

func(T AddNoteToNoteBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	AddMessage := AddJson{}

	json.Unmarshal(request.GetData(),&AddMessage)

	fmt.Println("Handle Addnotetobook 传来的信息：",AddMessage)

	Line := snet.SDBNote.Debug().Exec("call add_to_book(?,?)",AddMessage.bid,AddMessage.nid).RowsAffected

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