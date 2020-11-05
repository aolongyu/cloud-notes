package shandle

import (
	Data "data"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type UpdataNickName struct {
	snet.BaseRouter
}

func (U UpdataNickName) Handle(request isface.IRequest) {

	player := &Data.PlayerInfo{
		Roomid:   0,
		Uid:      0,
		Username: "",
	}
	proto.Unmarshal(request.GetData(), player)
	u := &Data.User_info{}
	snet.SDB.Debug().Table("user_info").Where("uid=?", player.Uid).Scan(u)
	oldnickname := u.User_name
	snet.SDB.Debug().Table("user_info").Where("uid=?", player.Uid).Update("user_name", player.Username).Scan(u)
	newnickname := u.User_name
	if oldnickname == newnickname {
		request.GetConnection().SendMesg(UPDATA_NICK_NAME_ACK, []byte("0")) //失败
	} else {
		request.GetConnection().SendMesg(UPDATA_NICK_NAME_ACK, []byte("1")) //成功
	}
}
