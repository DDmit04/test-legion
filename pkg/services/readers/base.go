package readers

import (
	"reflect"
	"test-legion/pkg/models"
	"test-legion/pkg/models/generator"
)

type PortReader interface {
	CanRead(dataType reflect.Type) bool
	Read(data any) (generator.ValueGenerator[int], error)
}

type HostReader interface {
	CanRead(dataType reflect.Type) bool
	Read(data any) (generator.ValueGenerator[models.DomainInfo], error)
}

type BaseReader struct {
	ReadType reflect.Type
}

func (r *BaseReader) CanRead(dataType reflect.Type) bool {
	return r.ReadType == dataType
}
