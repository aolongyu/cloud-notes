package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindByidDetail struct{
	snet.BaseRouter
}

type FindByidDetailJson struct{
	Uid int `json:"Uid"`
}

type FindByidDetailGorm struct{
	Name string	`gorm:"column:name"`
	Stats int	`gorm:"column:stats"`
	Id 	  int	`gorm:"column:id"`
	Phone string	`gorm:"column:phone"`
	Sex int	`gorm:"column:sex"`
	Uadd string	`gorm:"column:uadd"`
	Email string	`gorm:"column:email"`
}

func(T FindByidDetail)Handle(request isface.IRequest){
	conn := request.GetConnection()
	recvData := FindByidDetailJson{}

	json.Unmarshal(request.GetData(),&recvData)

	fmt.Println("Handle FindByidDetail 接收到消息",recvData)

	Data := FindByidDetailGorm{}

	snet.SDB.Debug().Raw("call findbyuid(?)",recvData.Uid).Scan(&Data)

	SendData,_ := json.Marshal(Data)
	conn.SendMesg([]byte(""),SendData)
}