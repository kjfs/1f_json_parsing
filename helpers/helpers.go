package helpers

import (
	"OneFootball_Zusammenfuehren/config"
	"fmt"
	"runtime/debug"
)

var conf *config.Config

// Sets up Config for helpers package.
func NewHelpers(a *config.Config) {
	conf = a
}

// Infolog prints all relevant information such as if in the JSON body some fields have no values. This data is created manually and passed as an arugment to this function.
func InfoLog(info string) {
	conf.InfoLog.Println("InfoLog: ", info)
}

// Errorlog stores all errors that may occur. It creates a debug strack trace of the go routine that calls it.
func ErrorLog(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	conf.ErrorLog.Println(trace)
}
