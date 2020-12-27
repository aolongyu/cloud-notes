package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type DeleteNote struct{
	snet.BaseRouter
}

type DeleteNoteJson struct{
	Sid int `json:"sid"`
	Note_id  int `json:"note_id"`
}

func(T DeleteNote)Handle(request isface.IRequest){
	conn := request.GetConnection()
	AddMessage := DeleteNoteJson{}

	json.Unmarshal(request.GetData(),&AddMessage)

	fmt.Println("Handle deletenote 传来的信息：",AddMessage)

	Line := snet.SDBNote.Debug().Exec("call delete_note(?,?)",AddMessage.Sid,AddMessage.Note_id).RowsAffected

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