package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type Register struct{
	snet.BaseRouter
}

type RegisterJson struct{
	Name string `json:"Name"`
	Password string `json:"Password"`
}

func(T Result) Handle(request isface.IRequest){
	conn := request.GetConnection()
	RegisterMessage := RegisterJson{}
	json.Unmarshal(request.GetData(),&RegisterMessage)

	res := Result{}
	fmt.Println("Handle Register 传来的信息：姓名",RegisterMessage.Name,"密码：",RegisterMessage.Password)

	snet.SDB.Debug().Raw("call register(?,?)",RegisterMessage.Name,RegisterMessage.Password).Scan(&res)

	fmt.Println("读取数据库的信息",res)
	returnres := Status{}
	//注册成功
	if res.Result == 0{
		returnres.Status = "1"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""), data)
	}else{
		returnres.Status = "1"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""), data)
	}
}