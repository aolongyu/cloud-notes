package shandle

import (
	"encoding/json"
	"fmt"
	"isface"
	"snet"
)

type ThumbsUpNoteBook struct{
	snet.BaseRouter
}

type ThumbsUpNoteBookJson struct{
	Nid int `json:"nid"`
	Bid int `json:"bid"`
}

func(T ThumbsUpNoteBook)Handle(request isface.IRequest){
	conn := request.GetConnection()
	RecvData := ThumbsUpNoteBookJson{}
	json.Unmarshal(request.GetData(),&RecvData)


	fmt.Println("Handle ThumbsUpNoteBook 传来的信息：",RecvData)

	Line := snet.SDBNote.Debug().Exec("call thumbs_up_notebook(?,?)",RecvData.Nid,RecvData.Bid).RowsAffected

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
