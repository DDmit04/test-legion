package readers

import (
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/models/input"
)

type Reader[T any] interface {
	CanRead(dataType input.Type) bool
	Read(data input.Model) (generator.ValueGenerator[T], error)
}

type BaseReader struct {
	ReadType input.Type
}

func (r *BaseReader) CanRead(inputType input.Type) bool {
	return r.ReadType == inputType
}
