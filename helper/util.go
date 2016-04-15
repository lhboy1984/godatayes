package helper

import "errors"

func IfErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func IfFailPanic(cond bool, info string) {
	if !cond {
		panic(errors.New(info))
	}
}
