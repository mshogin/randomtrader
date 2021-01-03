package strategy

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"plugin"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/logger"
)

type Processor func(ctx *bidcontext.BidContext) error

var plugins = map[string]Processor{}

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

		s, err := p.Lookup("ProcessContext")
		if err != nil {
			logger.Infof("plugin %q is not a strategy plugin: %w", pPath, err)
			continue
		}

		cp, ok := s.(func(ctx *bidcontext.BidContext) error)
		if !ok {
			logger.Infof("plugin %q does not implement ContextProcessor interface: %w", pPath, err)
			continue
		}

		plugins[f.Name()] = cp
	}
	return nil
}
