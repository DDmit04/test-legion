package port

import (
	"fmt"
	"reflect"
	"test-legion/pkg/models/generator"
	rdrs "test-legion/pkg/services/readers"
)

type ArgTypes interface {
	~[]int | ~string
}

var portsReaders = []rdrs.PortReader{
	NewListReader(),
	NewRangeReader(),
}

func ReadPorts[T ArgTypes](ports T) (generator.ValueGenerator[int], error) {
	for _, reader := range portsReaders {
		canRead := reader.CanRead(reflect.TypeOf(ports))
		if canRead {
			return reader.Read(ports)
		}
	}
	return nil, fmt.Errorf("not expected generator type")
}
