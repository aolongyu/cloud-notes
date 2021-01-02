package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ReportNoteNoteBook struct{
	snet.BaseRouter
}

type ReportNoteNoteBookJson struct{
	Uid int `json:"uid"`
	Bid int `json:"bid"`
}

func(T ReportNoteNoteBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := ReportNoteNoteBookJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle ReportNoteNoteBook 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call report_notebook(?,?)",RecvData.Uid,RecvData.Bid).RowsAffected

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
