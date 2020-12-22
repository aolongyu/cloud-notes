package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type CreateNoteBook struct{
	snet.BaseRouter
}
//用户名，笔记本名，笔记本简介，笔记本类型
//笔记本类型 0表示私有 1表示公开 2表示仅好友可见
type CreateNoteBookJson struct {
	Uid int `json:"Uid"`
	NoteBookName string `json:"NoteBookName"`
	NoteBookIntroduction string `json:"NoteBookIntroduction"`
	NoteBookType string `json:"NoteBookType"`
}

func (T CreateNoteBook)Handle(request isface.IRequest){
	conn := request.GetConnection()

	NoteJson := CreateNoteBookJson{}
	json.Unmarshal(request.GetData(),&NoteJson)

	fmt.Println("Handle CreateNoteBook 从客户端获得信息:",NoteJson)

	Line := snet.SDBNote.Debug().Exec("call create_notebook(?,?,?,?)",NoteJson.Uid,NoteJson.NoteBookName,NoteJson.NoteBookIntroduction,NoteJson.NoteBookType).RowsAffected
	fmt.Println("Line : ",Line)
	returnres := Status{}

	if Line > 0{
		returnres.Status = "1"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""),data)
	}else{
		returnres.Status = "0"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""),data)
	}

}
