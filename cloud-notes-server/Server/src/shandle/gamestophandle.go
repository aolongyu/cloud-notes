package shandle

import (
	Data "data"
	"fmt"
	"github.com/golang/protobuf/proto"
	"isface"
	"snet"
)

//游戏结束时候记录到数据库的内容

//ToDo加入到Handle中
type GameStopHandle struct {
	snet.BaseRouter
}

func (G *GameStopHandle)Handle(request isface.IRequest) {
	OGM := &Data.OverGameMessage{
		AllPlayerBehave: nil,
	}
	proto.Unmarshal(request.GetData(),OGM)

	//数组长度
	Len := len(OGM.AllPlayerBehave)

	cus := &Data.Account_info{}
	for i:=0; i<Len; i++{
		cus.UID = OGM.AllPlayerBehave[i].Uid
		snet.SDB.Table("account_info").Where("UID = ?",cus.UID).Scan(&cus)
		exp := cus.Account_exp - int32(2 * i) + 7
		snet.SDB.Exec("update account_info set account_exp = ? where Uid = ?",exp,cus.UID)

	}
	//snet.Logs.Info("该局信息更新到数据库")
	fmt.Println("该局信息更新到数据库")
}