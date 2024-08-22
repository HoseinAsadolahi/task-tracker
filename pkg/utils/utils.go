package utils

import (
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/statics"
)

func IfThenElse(cond bool, a, b any) any {
	if cond {
		return a
	}
	return b
}

func CheckError(err error, message string) bool {
	if err != nil {
		fmt.Println(statics.ErrorStyle.Render(message))
		return true
	}
	return false
}
