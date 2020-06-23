package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
)

var debugDisabled = true

func EnableDebug() func() {
	debugDisabled = false
	return DisableDebug
}

func DisableDebug() {
	debugDisabled = true
}

func Debug(format string, args ...interface{}) {
	if debugDisabled {
		return
	}
	_, fileName, fileLine, ok := runtime.Caller(1)
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d:", fileName, fileLine)
		p := strings.Split(s, "/")
		s = strings.Join(p[len(p)-2:], "/")
	} else {
		s = ""
	}
	fmt.Printf("DEBUG: "+s+" "+format+"\n", args...) // output for debug
}

func Error(format string, args ...interface{}) {
	fmt.Printf("ERROR: "+format+"\n", args...) // output for debug
}

func Info(format string, args ...interface{}) {
	fmt.Printf("INFO: "+format+"\n", args...) // output for debug
}

func ProcessContext(ctx *bidcontext.BidContext) error {
	buf, err := json.MarshalIndent(ctx, "", "    ")
	if err != nil {
		err = fmt.Errorf("cannot dump context: %w", err)
		Error(err.Error())
		return err
	}
	Info(string(buf))
	return nil
}
