package port

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"test-legion/pkg/models/exceptions"
	"test-legion/pkg/models/generator"
	"test-legion/pkg/services/readers"
)

const rangeSeparator = "-"

type RangeReader struct {
	BasePortReader
}

func NewRangeReader() *RangeReader {
	return &RangeReader{
		BasePortReader: BasePortReader{
			BaseReader: readers.BaseReader{
				ReadType: reflect.TypeOf(""),
			},
		},
	}
}

func (r *RangeReader) Read(data any) (generator.ValueGenerator[int], error) {
	rang := data.(string)
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
	return generator.NewPortsRange(start, end), nil
}

func (r *RangeReader) parseRange(rang string) (int, int, error) {
	parts := strings.Split(rang, rangeSeparator)
	if len(parts) != 2 {
		return 0, 0, exceptions.UnexpectedRangeFormat(fmt.Sprintf("invalid range - \"%s\"", rang))
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, exceptions.InvalidPortValue(fmt.Sprintf("port not a number - \"%s\"", parts[0]))
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, exceptions.InvalidPortValue(fmt.Sprintf("port not a number - \"%s\"", parts[1]))
	}
	return start, end, nil
}
