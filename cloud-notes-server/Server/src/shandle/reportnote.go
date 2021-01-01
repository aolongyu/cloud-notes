package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ReportNote struct{
	snet.BaseRouter
}

type ReportNoteJson struct{
	uid int `json:"uid"`
	nid int `json:"nid"`
}

func(T ReportNote)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := ReportNoteJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle ReportNote 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call report_note(?,?)",RecvData.uid,RecvData.nid).RowsAffected

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
