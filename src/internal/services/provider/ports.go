package provider

import (
	"fmt"

	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	rdrs "github.com/DDmit04/test-legion/src/internal/services/readers"
	port2 "github.com/DDmit04/test-legion/src/internal/services/readers/port"
	"github.com/DDmit04/test-legion/src/models/input"
)

var portsReaders = []rdrs.Reader[int]{
	port2.NewListReader(),
	port2.NewRangeReader(),
}

type PortsProvider struct {
	Readers []rdrs.Reader[int]
}

func NewPortsProvider() *PortsProvider {
	return &PortsProvider{
		Readers: []rdrs.Reader[int]{
			port2.NewListReader(),
			port2.NewRangeReader(),
		},
	}
}

func (r *PortsProvider) ReadData(ports input.Model) (generator.ValueGenerator[int], error) {
	for _, reader := range portsReaders {
		canRead := reader.CanRead(ports.InputType())
		if canRead {
			return reader.Read(ports)
		}
	}
	msg := fmt.Sprintf("ports reader with name: \"%s\" not exists", ports.InputType())
	return nil, exceptions.ReaderNotFround(msg)
}
