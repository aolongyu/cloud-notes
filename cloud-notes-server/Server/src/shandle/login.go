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
	//Mark int `gorm:"column:Mark"`
}
type Status struct{
	Status string  //返回的状态
}
type LoginStatus struct{
	Status string
	Uid string
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
	returnres := LoginStatus{
		Status: "0",
		Uid:    "111",
	}
	//登录成功
	if res.Result > 0 {
		if res.Result == 1{
			returnres.Status = "1"
		}else if res.Result == 2{
			returnres.Status = "2"
		}
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""), data)
	}else{//登录失败.+封禁+无此用户
		returnres.Status = "0"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""), data)
	}
}
