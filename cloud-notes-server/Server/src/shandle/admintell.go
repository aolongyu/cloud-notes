package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type AdminTell struct{
	snet.BaseRouter
}

type AdminTellJson struct{
	Utext string `json:"utext"`
}


func(T AdminTell)Handle(request isface.IRequest){
	conn := request.GetConnection()
	AddMessage := AdminTellJson{}

	json.Unmarshal(request.GetData(),&AddMessage)

	fmt.Println("Handle AdminTellJson 传来的信息：",AddMessage)

	Line := snet.SDB.Debug().Exec("call admin_tell(?)",AddMessage.Utext).RowsAffected

	returnres := Status{}
	if(Line > 0){
		returnres.Status = "1"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""),data)
	}else{
		returnres.Status = "0"
		data,_ := json.Marshal(returnres)
		conn.SendMesg([]byte(""),data)
	}
}