package shandle

import (
	Data "data"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type UpdataPass struct {
	snet.BaseRouter
}

func (U UpdataPass) Handle(request isface.IRequest) {
	cus := &Data.Customer{}
	proto.Unmarshal(request.GetData(), cus)
	u := &Data.User_info{}
	snet.SDB.Debug().Table("user_info").Where("account_name=?", cus.User).Scan(u)
	oldpass := u.Account_password
	snet.SDB.Debug().Table("user_info").Where("account_name=?", cus.User).Update("account_password", cus.Pass).Scan(u)
	newpass := u.Account_password
	if oldpass == newpass {
		request.GetConnection().SendMesg(UPDATA_PASS_ACK, []byte("0")) //失败
	} else {
		request.GetConnection().SendMesg(UPDATA_PASS_ACK, []byte("1")) //成功
	}
}
