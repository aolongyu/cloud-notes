package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type Login struct{
	snet.BaseRouter
}
type LoginUserJson struct{
	Name string `json:"Name"`
	Password string `json:"Password"`
}
type Result struct{
	Result int `gorm:"column:Result"`
}
func(T Login) Handle(request isface.IRequest){
	conn := request.GetConnection()
	Loginmessage := LoginUserJson{}

	json.Unmarshal(request.GetData(),&Loginmessage)
	//登录失败
	res := Result{}
	fmt.Println("Handle Login   传来的信息:姓名",Loginmessage.Name,"密码",Loginmessage.Password)

	snet.SDB.Debug().Raw("call login(?,?)",Loginmessage.Name,Loginmessage.Password).Scan(&res)

	fmt.Println("读取的数据库内容：",res)

	if res.Result > 0 {
		conn.SendMesg([]byte("loginack"), []byte("ok"))
	}else{
		conn.SendMesg([]byte("loginack"), []byte("no"))
	}
}
