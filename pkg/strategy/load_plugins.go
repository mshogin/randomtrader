package strategy

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"plugin"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
)

const processContextCallbackName = "ProcessContext"
const runRoutineCallbackName = "RunRoutine"

type Processor func(ctx *bidcontext.BidContext) error

var pluginContextProcessors = map[string]Processor{}
var pluginRoutines = map[string]func(){}

// Init ...
func Init() error {
	files, err := ioutil.ReadDir(config.GetPluginsDir())
	if err != nil {
		return fmt.Errorf("cannot read plugins directory: %w", err)
	}

	for _, f := range files {
		pPath := filepath.Join(config.GetPluginsDir(), f.Name())
		p, err := plugin.Open(pPath)
		if err != nil {
			return fmt.Errorf("%q is not a go plugin: %w", pPath, err)
		}

		s, err := p.Lookup(processContextCallbackName)
		if err == nil {
			cp, ok := s.(func(ctx *bidcontext.BidContext) error)
			if ok {
				pluginContextProcessors[f.Name()] = cp
			}
		}

		r, err := p.Lookup(runRoutineCallbackName)
		if err == nil {
			routine, ok := r.(func())
			if ok {
				go routine()
			}
		}
	}
	return nil
}
