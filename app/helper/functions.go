package helper

import (
	"runtime"
	"strings"
)

func GetCurFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	name = strings.Split(name, ".")[2]
	return strings.ToLower(name)
}
