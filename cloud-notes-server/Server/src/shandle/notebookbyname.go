package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindNoteBookByUName struct{
	snet.BaseRouter
}
//用户名
type FindNoteBookByNameJson struct{
	name string 	`json:"uname"`
}

func(T FindNoteBookByUName)Handle(request isface.IRequest){
	conn := request.GetConnection()
	username := FindNoteBookByNameJson{}
	json.Unmarshal(request.GetData(),&username)

	fmt.Println("Handle FindNoteBookByNameJson 传来的信息:",username)
	data := make([]NoteFindByUser,0)
	snet.SDBNote.Debug().Raw("call user_notebook_byname(?)",username.name).Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}
