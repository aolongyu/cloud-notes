package main

import (
	"snet"
	"shandle"
)
func main(){
	s:= snet.NewServer()

	//在这里添加handle键值对
	shandle.AddHandleInit(s)

	s.Serve()

}
