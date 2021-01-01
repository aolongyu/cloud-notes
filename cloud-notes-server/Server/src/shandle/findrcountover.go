package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindRountOver struct{
	snet.BaseRouter
}

type FindRountOverGorm struct{
	Id int	`gorm:"column:"id"`
	Text string	`gorm:"column:"text"`
	Report int	`gorm:"column:"report"`
	Introduction string	`gorm:"column:"introduction"`
}

func(T FindRountOver)Handle(request isface.IRequest){
	conn := request.GetConnection()

	fmt.Println("Handle FindRountOver 传来的信息：")

	data := make([]FindRountOverGorm,0)
	snet.SDBNote.Debug().Raw("call findrcount_over10()").Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}
