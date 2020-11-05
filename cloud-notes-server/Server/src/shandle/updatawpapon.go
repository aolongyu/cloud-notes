package shandle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"isface"
	"snet"
)

type UpdataWpapon struct {
	snet.BaseRouter
}

func (U UpdataWpapon)Handle(request isface.IRequest){
	uid := request.GetConnection().GetConnID()
	bytesBuffer := bytes.NewBuffer(request.GetData())
	var x int32
	binary.Read(bytesBuffer,binary.LittleEndian,&x)

	snet.SDB.Debug().Exec("update account_info set weapon_id = ? where UID = ?",x,uid)

	request.GetConnection().SendMesg(UPDATA_WPAPON_ACK,[]byte("1") )
	fmt.Println("用户为",uid,"的用户武器信息更新到数据库中了，绑定武器为",x)
}