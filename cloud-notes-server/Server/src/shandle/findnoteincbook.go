package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindAllCollectInBook struct{
	snet.BaseRouter
}

type FindAllCollectInBookJson struct{
	Id int	`json:"bid"`
}

type FindAllCollectInBookGorm struct{
	Name int `gorm:"column:name"`
	Introduction string `gorm:"column:introduction"`
	Type int `gorm:"column:type"`
	Id string `gorm:"column:id"`
}

func(T FindAllCollectInBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	Ida := FindAllCollectInBookJson{}
	json.Unmarshal(request.GetData(),&Ida)

	fmt.Println("Handle FindAllCollectBook   传来的信息:  ",Ida.Id)
	data := make([]FindAllCollectInBookGorm,0)
	snet.SDBNote.Debug().Raw("call find_note_in_cbook(?)",Ida.Id).Scan(&data)

	SendData,_ := json.Marshal(data)
	conn.SendMesg([]byte(""),SendData)
}
