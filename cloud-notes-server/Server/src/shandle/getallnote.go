package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type GetAllNote struct{
	snet.BaseRouter
}


type GetAllNoteGorm struct{
	Id int   `gorm:"column:note_id"`
	Name string	`gorm:"column:note_name"`
	Uid string	`gorm:"column:uid"`
	Uname string `gorm:"column:uname"`
	Thumbs int	`gorm:"column:thumbs"`
}


func(T GetAllNote)Handle(request isface.IRequest){
	conn := request.GetConnection()

	fmt.Println("Handle GetAllNote   传来的信息:",)
	data := make([]GetAllNoteGorm,0)
	snet.SDBNote.Debug().Raw("call get_all_note()").Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}