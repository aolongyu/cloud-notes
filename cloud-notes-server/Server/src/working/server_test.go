package main

import (
	"shandle"
	"snet"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	b.StartTimer()
	s := snet.NewServer()

	//在这里添加handle键值对
	shandle.AddHandleInit(s)

	s.Serve()
	b.StopTimer()
}
