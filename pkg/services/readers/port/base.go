package port

import (
	"test-legion/pkg/models/exceptions"
	"test-legion/pkg/services/readers"
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
	if port > 65535 {
		return exceptions.TooHighPortValue()
	}
	return nil
}
