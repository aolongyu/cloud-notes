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
	Uid int `json:"uid"`
	Nid int `json:"nid"`
}

type StatusReoprt struct{
	Result int `gorm:"column:@i"`
}

func(T ReportNote)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := ReportNoteJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle ReportNote 传来的信息：",RecvData)
	Data := StatusReoprt{}

	snet.SDBNote.Debug().Exec("call report_note(?,?)",RecvData.Uid,RecvData.Nid).Scan(&Data)

	res := Status{}
	if Data.Result  == 0{
		res.Status = "1"
		SendData,_ := json.Marshal(res)
		conn.SendMesg([]byte(""),SendData)
	}else{
		res.Status = "0"
		SendData,_ := json.Marshal(res)
		conn.SendMesg([]byte(""),SendData)
	}
}
