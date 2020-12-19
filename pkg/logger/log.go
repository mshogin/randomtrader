package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
)

var debugDisabled = true

// EnableDebug ...
func EnableDebug() func() {
	debugDisabled = false
	return DisableDebug
}

// DisableDebug ...
func DisableDebug() {
	debugDisabled = true
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
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
	fmt.Printf("DEBUG: "+s+" "+format+"\n", args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	fmt.Printf("ERROR: "+format+"\n", args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf("FATAL: "+format+"\n", args...))
}

// Infof ...
func Infof(format string, args ...interface{}) {
	fmt.Printf("INFO: "+format+"\n", args...)
}

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	buf, err := json.MarshalIndent(ctx, "", "    ")
	if err != nil {
		err = fmt.Errorf("cannot dump context: %w", err)
		Errorf(err.Error())
		return err
	}
	Infof(string(buf))
	return nil
}
