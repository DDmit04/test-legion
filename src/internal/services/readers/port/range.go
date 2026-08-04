package port

import (
	"fmt"
	"strconv"
	"strings"

	exceptions2 "github.com/DDmit04/test-legion/src/internal/models/exceptions"
	generator2 "github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
	"github.com/DDmit04/test-legion/src/models/input"
)

const rangeSeparator = "-"

type RangeReader struct {
	BasePortReader
}

func NewRangeReader() *RangeReader {
	return &RangeReader{
		BasePortReader: BasePortReader{
			BaseReader: readers.BaseReader{
				ReadType: input.PortsRangeInputType,
			},
		},
	}
}

func (r *RangeReader) Read(data input.Model) (generator2.ValueGenerator[int], error) {

	rang := ""
	if dataStr, ok := data.Value().(string); ok {
		rang = dataStr
	} else {
		msg := fmt.Sprintf("expected string, got %T", data)
		return nil, exceptions2.InputDataMisMatch(msg)
	}

	start, end, err := r.parseRange(rang)
	if err != nil {
		return nil, err
	}
	for _, port := range []int{start, end} {
		validateErr := r.validatePort(port)
		if validateErr != nil {
			validateErr = validateErr.WithMsg(fmt.Sprintf("range - \"%s\"", rang))
			return nil, validateErr
		}
	}
	return generator2.NewPortsRange(start, end), nil
}

func (r *RangeReader) parseRange(rang string) (int, int, error) {
	parts := strings.Split(rang, rangeSeparator)
	if len(parts) != 2 {
		return 0, 0, exceptions2.UnexpectedRangeFormat(fmt.Sprintf("invalid range - \"%s\"", rang))
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, exceptions2.InvalidPortValue(fmt.Sprintf("port not a number - \"%s\"", parts[0]))
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, exceptions2.InvalidPortValue(fmt.Sprintf("port not a number - \"%s\"", parts[1]))
	}
	if start > end {
		start, end = end, start
	}
	return start, end, nil
}
