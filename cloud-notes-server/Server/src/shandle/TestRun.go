package shandle

import (
	"fmt"
	"isface"
	"snet"
)

type TestRun struct {
	snet.BaseRouter
}

func (S *TestRun)Handle(request isface.IRequest){
	fmt.Println("删除房间")
	ALLROOM := snet.RoomMgr.AllRoom

	for k,v := range ALLROOM{
		v.Lock()
		delete(ALLROOM,k)
		v.Unlock()
	}
}