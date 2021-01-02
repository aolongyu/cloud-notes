package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindAllCollectBook struct{
	snet.BaseRouter
}

type FindAllCollectBookJson struct{
	Id int	`json:"uid"`
}

type FindAllCollectBookGorm struct{
	Id int `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Introduction string `gorm:"column:introduction"`
	Type int `gorm:"column:type"`
}

func(T FindAllCollectBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	Ida := FindAllCollectBookJson{}
	json.Unmarshal(request.GetData(),&Ida)

	fmt.Println("Handle FindAllCollectBook   传来的信息:  ",Ida.Id)
	data := make([]FindAllCollectBookGorm,0)
	snet.SDBNote.Debug().Raw("call find_All_Collectbook(?)",Ida.Id).Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}
