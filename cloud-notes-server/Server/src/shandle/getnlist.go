package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)
//用用户名获取笔记本列表
type GetNlist struct{
	snet.BaseRouter
}

type UserNameJson struct{
	Uid int `json:"Uid"`
}


type NoteList struct{
	Id int   `gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Introduction string	`gorm:"column:introduction"`
	ThumbsUp int	`gorm:"column:thumbs_up"`
}

func(T GetNlist)Handle(request isface.IRequest){
	conn := request.GetConnection()
	UserName := UserNameJson{}

	json.Unmarshal(request.GetData(),&UserName)

	fmt.Println("Handle GetNList 传来的信息：",UserName)
	data := make([]NoteList,0)
	snet.SDBNote.Debug().Raw("call user_notebook(?)",UserName.Uid).Scan(&data)

	Sendata,_ := json.Marshal(data)

	conn.SendMesg([]byte(""),Sendata)
}
