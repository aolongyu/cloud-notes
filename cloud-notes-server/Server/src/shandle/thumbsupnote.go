package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ThumbsUpNote struct{
	snet.BaseRouter
}

type ThumbsUpNoteJson struct{
	Nid int `json:"uid"`
	Bid int `json:"nid"`
}

func(T ThumbsUpNote)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := ThumbsUpNoteJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle ThumbsUpNote 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call thumbs_up_note(?,?)",RecvData.Nid,RecvData.Bid).RowsAffected

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
