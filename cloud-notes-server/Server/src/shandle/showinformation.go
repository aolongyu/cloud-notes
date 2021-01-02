package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ShowInformation struct{
	snet.BaseRouter
}

type ShowInformationJson struct{
	PageNo int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}


type ShowInformationGorm struct{
	Id int   `gorm:"column:id"`
	Text string	`gorm:"column:text"`
	Introduction string	`gorm:"column:introduction"`
	ModifiedTime JSONTime `gorm:"column:modified_time"`
}

func(T ShowInformation)Handle(request isface.IRequest){
	conn := request.GetConnection()
	UserName := ShowInformationJson{}

	json.Unmarshal(request.GetData(),&UserName)

	fmt.Println("Handle ShowInformation 传来的信息：",UserName)
	data := make([]ShowInformationGorm,0)
	snet.SDBNote.Debug().Raw("call show_information(?,?)",UserName.PageNo,UserName.PageSize).Scan(&data)

	Sendata,_ := json.Marshal(data)

	conn.SendMesg([]byte(""),Sendata)
}

