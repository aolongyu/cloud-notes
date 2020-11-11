package shandle

import (
	"isface"
	"snet"
)

type Nofound struct {
	snet.BaseRouter
}

func (this *Nofound)Handle (request isface.IRequest){
	wohah:= []byte("wori")
	request.GetConnection().SendMesg(3,wohah)
}