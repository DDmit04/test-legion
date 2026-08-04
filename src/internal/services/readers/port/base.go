package port

import (
	exceptions2 "github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
)

const (
	maxPort = 65535
)

type BasePortReader struct {
	readers.BaseReader
}

func (p *BasePortReader) validatePort(port int) *exceptions2.ModuleException {
	if port < 0 {
		return exceptions2.NegativePortValue()
	}
	if port == 0 {
		return exceptions2.ZeroPortValue()
	}
	if port > maxPort {
		return exceptions2.TooHighPortValue()
	}
	return nil
}
