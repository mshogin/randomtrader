package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveOrderBookLog(t *testing.T) {
	s := assert.New(t)

	gceClientOrig := SwapGCEClient(GetGCETestClient())
	defer func() { SwapGCEClient(gceClientOrig) }()

	s.NoError(SaveOrderBookLog("/some/path"))
}
