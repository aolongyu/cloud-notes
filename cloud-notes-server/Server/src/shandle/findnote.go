package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindNoteByUserName struct{
	snet.BaseRouter
}

type NoteFindByUser struct{
	Id int   `gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Introduction string	`gorm:"column:introduction"`
	Text string `gorm:"column:text"`
	ThumbsUp int	`gorm:"column:thumbs_up"`
}


func(T FindNoteByUserName)Handle(request isface.IRequest){
	conn := request.GetConnection()
	username := UserNameJson{}
	json.Unmarshal(request.GetData(),&username)

	fmt.Println("Handle FindNotaeByUserName   传来的信息:姓名",username.Uid)
	data := make([]NoteFindByUser,0)
	snet.SDBNote.Debug().Raw("call user_note(?)",username.Uid).Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}