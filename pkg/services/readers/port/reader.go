package port

import (
	"fmt"

	"github.com/DDmit04/test-legion/pkg/models/generator"
	"github.com/DDmit04/test-legion/pkg/models/input"
	rdrs "github.com/DDmit04/test-legion/pkg/services/readers"
)

var portsReaders = []rdrs.Reader[int]{
	NewListReader(),
	NewRangeReader(),
}

func ReadPorts(ports input.Model) (generator.ValueGenerator[int], error) {
	for _, reader := range portsReaders {
		canRead := reader.CanRead(ports.InputType())
		if canRead {
			return reader.Read(ports)
		}
	}
	return nil, fmt.Errorf("not expected generator type")
}
