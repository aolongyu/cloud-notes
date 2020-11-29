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

	if res.result == 0{
		conn.SendMesg([]byte("regiack"),[]byte("ok"))
	}else{
		conn.SendMesg([]byte("regiack"),[]byte("no"))
	}
}