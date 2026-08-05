package port

import (
	"fmt"

	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
	"github.com/DDmit04/test-legion/src/internal/tools"
	"github.com/DDmit04/test-legion/src/models/input"
)

type ListReader struct {
	BasePortReader
}

func NewListReader() *ListReader {
	return &ListReader{
		BasePortReader: BasePortReader{
			BaseReader: readers.BaseReader{
				ReadType: input.PortsListInputType,
			},
		},
	}
}

func (r *ListReader) Read(data input.Model) (generator.ValueGenerator[int], error) {
	prts := make([]int, 0)
	if dataInts, ok := data.Value().([]int); ok {
		prts = tools.UniqueIntsList(dataInts)
	} else {
		msg := fmt.Sprintf("expected []int, got %T", data)
		return nil, exceptions.InputDataMisMatch(msg)
	}
	for _, port := range prts {
		validateErr := r.validatePort(port)
		if validateErr != nil {
			validateErr = validateErr.WithMsg(fmt.Sprintf("error port - \"%d\"", port))
		}
	}
	return generator.NewPortsList(prts), nil
}
