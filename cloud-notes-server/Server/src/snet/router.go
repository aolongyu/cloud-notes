package snet

import "isface"
//实现router时候，先根据需要对其重写
type BaseRouter struct {

}

func (b BaseRouter)Handle(request isface.IRequest){

}
