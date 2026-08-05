package port

import (
	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
)

const (
	maxPort = 65535
)

type BasePortReader struct {
	readers.BaseReader
}

func (p *BasePortReader) validatePort(port int) *exceptions.ModuleException {
	if port < 0 {
		return exceptions.NegativePortValue()
	}
	if port == 0 {
		return exceptions.ZeroPortValue()
	}
	if port > maxPort {
		return exceptions.TooHighPortValue()
	}
	return nil
}
