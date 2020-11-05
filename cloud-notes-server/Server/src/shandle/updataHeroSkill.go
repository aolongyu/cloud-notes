package shandle

import (
	"bytes"
	"encoding/binary"
	"isface"
	"snet"
)

type UpdataHeroSkill struct {
	snet.BaseRouter
}

//由房主预处理信息给服务器，并且发送到服务器上。
func (U *UpdataHeroSkill) Handle(request isface.IRequest) {
	Uid := request.GetConnection().GetConnID()
	bytedata := bytes.NewBuffer(request.GetData())
	var num int32
	err := binary.Read(bytedata, binary.LittleEndian, &num)
	if err != nil {
		println(err)
	}
	skill := num % 10
	hero := num / 10
	snet.SDB.Debug().Table("account_info").Where("UID =?", Uid).Update("choosen_hero_id", hero)
	snet.SDB.Debug().Table("hero").Where("UID =? AND hero_id=?", Uid, hero).Update("hero_skill", skill)
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, 1)
	request.GetConnection().SendMesg(UPDATA_HEROSKILL_ACK, buf)

}
