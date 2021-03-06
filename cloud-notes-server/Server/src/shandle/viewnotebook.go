package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ViewNote struct{
	snet.BaseRouter
}
//Id号笔记Id
type ViewNoteJson struct{
	Id int `json:"id"`
}

type ViewNoteGorm struct {
	Id int `gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Introduction string	`gorm:"column:introduction"`
	Text string `gorm:"column:text"`
	ThumbsUp int	`gorm:"column:thumbs_up"`
}

func (T ViewNote)Handle(request isface.IRequest){
	conn := request.GetConnection()
	ViewNoteMessage := ViewNoteJson{}
	json.Unmarshal(request.GetData(),&ViewNoteMessage)

	fmt.Println("Handle ViewNote   传来的信息:",ViewNoteMessage)

	data := make([]ViewNoteGorm,0)
	snet.SDBNote.Debug().Raw("call note_inform(?)",ViewNoteMessage.Id).Scan(&data)
	SendData,_ := json.Marshal(data)

	conn.SendMesg([]byte(""),SendData)
}