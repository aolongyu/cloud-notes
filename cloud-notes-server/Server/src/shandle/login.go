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

func(T Login) Handle(request isface.IRequest){
	conn := request.GetConnection()
	Loginmessage := LoginUserJson{}

	json.Unmarshal(request.GetData(),&Loginmessage)

	fmt.Println("Handle Login   传来的信息:姓名",Loginmessage.Name,"密码",Loginmessage.Password)

	if Loginmessage.Name == "aolyu@qq.com" && Loginmessage.Password == "123" {
		conn.SendMesg([]byte("loginack"), []byte("ok"))
	}else{
		conn.SendMesg([]byte("loginack"), []byte("no"))
	}
}
