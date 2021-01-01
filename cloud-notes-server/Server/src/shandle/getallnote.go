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
	Id int   `gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Introduction string	`gorm:"column:introduction"`
	Text string `gorm:"column:text"`
	ThumbsUp int	`gorm:"column:thumbs_up"`
}


func(T GetAllNote)Handle(request isface.IRequest){
	conn := request.GetConnection()

	fmt.Println("Handle GetAllNote   传来的信息:",)
	data := make([]NoteFindByUser,0)
	snet.SDBNote.Debug().Raw("call get_all_note()").Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}