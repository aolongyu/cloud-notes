package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

//创建笔记

type CreateNote struct{
	snet.BaseRouter
}

type CreateNoteJson struct{
	Uid int `json:"Uid"`
	NoteName string	`json:"NoteName"`
	NoteIntroduction string	`json:"NoteIntroduction"`
	NodeType int	`json:"NodeType"`
	NodeText string	`json:"NodeText"`
}
func (T CreateNote) Handle(request isface.IRequest){
	conn := request.GetConnection()
	CreateNoteMesg := CreateNoteJson{}

	json.Unmarshal(request.GetData(),&CreateNoteMesg)

	fmt.Println("Handle createnote 传来的信息：",CreateNoteMesg)

	Line := snet.SDBNote.Debug().Exec("create_note(?,?,?,?,?)",CreateNoteMesg.Uid,CreateNoteMesg.NoteName,CreateNoteMesg.NoteIntroduction,CreateNoteMesg.NodeType,CreateNoteMesg.NodeText).RowsAffected

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