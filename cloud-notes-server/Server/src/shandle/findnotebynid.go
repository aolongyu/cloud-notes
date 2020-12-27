package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type GetNoteListById struct{
	snet.BaseRouter
}
//笔记本ID
type NoteListJson struct{
	id int	`json:"id"`
}

func(T GetNoteListById)Handle(request isface.IRequest){
	conn := request.GetConnection()
	username := NoteListJson{}
	json.Unmarshal(request.GetData(),&username)

	fmt.Println("Handle GetNoteListById   传来的信息:",username)
	data := make([]NoteFindByUser,0)
	snet.SDBNote.Debug().Raw("call notebook_note(?)",username.id).Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}