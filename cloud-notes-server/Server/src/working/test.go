package main

import (
	"fmt"
)

func main() {
	NeedExp := make([]int32,36)
	NeedExp[1] = 0
	NeedExp[2] = 6
	NeedExp[3] = 14
	for i := 4;i<=19;i++{
		NeedExp[i] = NeedExp[i-1]+7
	}
	temp := 14
	for i := 20;i<=35;i++{
		NeedExp[i] = NeedExp[i-1] + int32(temp)
		temp++
	}
	fmt.Println(NeedExp)
}