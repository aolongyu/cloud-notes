package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type UpdateInformation struct{
	snet.BaseRouter
}

type UpdateInformationJson struct{
	Id int `json:"id"`
	Imf string `json:"imf"`
}

func(T UpdateInformation)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := UpdateInformationJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle NoteBookcloseJson 传来的信息：",RecvData)

	Line := snet.SDB.Debug().Exec("call update_information(?,?)",RecvData.Id,RecvData.Imf).RowsAffected

	fmt.Println("Line : ",Line)
	res := Status{}
	if Line > 0{
		res.Status = "1"
		SendData,_ := json.Marshal(res)
		conn.SendMesg([]byte(""),SendData)
	}else{
		res.Status = "0"
		SendData,_ := json.Marshal(res)
		conn.SendMesg([]byte(""),SendData)
	}
}

