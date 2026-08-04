package host

import (
	"fmt"
	"reflect"
	"test-legion/pkg/models"
	"test-legion/pkg/models/generator"
	rdrs "test-legion/pkg/services/readers"
)

type ArgsType interface {
	~[]string
}

var hostReaders = []rdrs.HostReader{
	NewListReader(),
}

func ReadHosts[T ArgsType](ports T) (generator.ValueGenerator[models.DomainInfo], error) {
	for _, reader := range hostReaders {
		canRead := reader.CanRead(reflect.TypeOf(ports))
		if canRead {
			return reader.Read(ports)
		}
	}
	return nil, fmt.Errorf("not expected generator type")
}
