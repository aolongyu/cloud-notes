package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type FindUser struct{
	snet.BaseRouter
}

//管理员查找用户
//页码，页面大小，i是查找方法，sx是关键词
//目前只有i=1的一种查找方法，后面可能要做热度查找
type FindUserJson struct{
	PageNo int `json:PageNo`
	PageSize int `json:"Pagesize"`
	I int `json:"I"`
	Sx string `json:"SX"`
}

type FindUserGorm struct{
	Customer_id int `gorm:"column:customer_id"`
	LoginName string `gorm:"column:login_name"`
	Password string `gorm:"column:password"`
	UserStats int `gorm:"column:user_stats"`
	ModifiedTime string `gorm:"column:modified_time"`
	CustomerLogincol string `gorm:"column:customer_logincol"`
}
func(T FindUser)Handle(request isface.IRequest){
	conn := request.GetConnection()
	recvData := FindUserJson{}

	json.Unmarshal(request.GetData(),&recvData)

	fmt.Println("find Handle 从客户端接收到消息：",recvData)

	Data := make([]FindUserGorm,0)

	snet.SDB.Debug().Raw("call find(?,?,?,?)",recvData.PageNo,recvData.PageSize,recvData.I,recvData.Sx).Scan(&Data)

	SendData,_ := json.Marshal(Data)

	conn.SendMesg([]byte(""),SendData)
}

