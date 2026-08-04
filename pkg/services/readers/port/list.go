package port

import (
	"fmt"
	"reflect"
	"test-legion/pkg/models/generator"
	"test-legion/pkg/services/readers"
	"test-legion/tools"
)

type ListReader struct {
	BasePortReader
}

func NewListReader() *ListReader {
	return &ListReader{
		BasePortReader: BasePortReader{
			BaseReader: readers.BaseReader{
				ReadType: reflect.TypeOf(make([]int, 0)),
			},
		},
	}
}

func (r *ListReader) Read(data any) (generator.ValueGenerator[int], error) {
	prts := tools.UniqueIntsList(data.([]int))
	for _, port := range prts {
		validateErr := r.validatePort(port)
		if validateErr != nil {
			validateErr = validateErr.WithMsg(fmt.Sprintf("error port - \"%s\"", port))
		}
	}
	return generator.NewPortsList(prts), nil
}
