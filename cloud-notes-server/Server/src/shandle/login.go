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
	result int32
}
func(T Login) Handle(request isface.IRequest){
	conn := request.GetConnection()
	Loginmessage := LoginUserJson{}

	json.Unmarshal(request.GetData(),&Loginmessage)
	//登录失败
	res := &Result{result:0}
	fmt.Println("Handle Login   传来的信息:姓名",Loginmessage.Name,"密码",Loginmessage.Password)

	snet.SDB.Debug().Raw("call login(?,?)",Loginmessage.Name,Loginmessage.Password).Scan(res)

	if res.result == 0 {
		conn.SendMesg([]byte("loginack"), []byte("ok"))
	}else{
		conn.SendMesg([]byte("loginack"), []byte("no"))
	}
}
