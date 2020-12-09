package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveOrderBookLog(t *testing.T) {
	s := assert.New(t)

	GetGCEClientOrig := GetGCEClient
	defer func() { GetGCEClient = GetGCEClientOrig }()
	GetGCEClient = GetGCETestClient

	s.NoError(SaveOrderBookLog("/some/path"))
}
