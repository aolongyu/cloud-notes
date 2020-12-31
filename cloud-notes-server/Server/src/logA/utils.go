package logA

import (
	"path"
	"runtime"
)

//工具函数
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	fileName = path.Base(fileName)
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}
