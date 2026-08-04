package provider

import (
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/models/input"
)

type Provider[T any] interface {
	ReadData(data input.Model) (generator.ValueGenerator[T], error)
}
