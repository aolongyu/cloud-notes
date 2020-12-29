package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindUserName struct{
	snet.BaseRouter
}

//管理员查找用户
//页码，页面大小，i是查找方法，sx是关键词
//目前只有i=1的一种查找方法，后面可能要做热度查找
type FindUserNameJson struct{
	PageNo int `json:PageNo`
	PageSize int `json:"Pagesize"`
	Sx string `json:"SX"`
}

type FindUserNameGorm struct{
	Name string `gorm:"column:name"`
	Stats int `gorm:"column:stats"`
	ModifiedTime string `gorm:"column:modified_time"`
}


func(T FindUserName)Handle(request isface.IRequest){
	conn := request.GetConnection()
	recvData := FindUserNameJson{}

	json.Unmarshal(request.GetData(),&recvData)

	fmt.Println("findusername Handle 从客户端接收到消息：",recvData)

	Data := make([]FindUserGorm,0)

	snet.SDB.Debug().Exec("call find_username(?,?,?)",recvData.PageNo,recvData.PageSize,recvData.Sx).Scan(&Data)

	SendData,_ := json.Marshal(Data)

	conn.SendMesg([]byte(""),SendData)
}