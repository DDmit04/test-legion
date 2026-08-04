package input

import "reflect"

const (
	PortsListInputType  Type = "ports_array"
	PortsRangeInputType Type = "ports_range"
)

type PortsListInput struct {
	BaseInputModel
}

func CreatePortsListInput(ports []int) (*PortsListInput, error) {
	res := &PortsListInput{
		BaseInputModel: BaseInputModel{
			value:        ports,
			inputType:    PortsListInputType,
			expectedType: reflect.TypeOf([]int{}),
		},
	}
	err := res.CheckType()
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PortsRangeInput struct {
	BaseInputModel
}

func CreatePortsRangeInput(portsRange string) (*PortsRangeInput, error) {
	res := &PortsRangeInput{
		BaseInputModel: BaseInputModel{
			value:        portsRange,
			inputType:    PortsRangeInputType,
			expectedType: reflect.TypeOf(""),
		},
	}
	err := res.CheckType()
	if err != nil {
		return nil, err
	}
	return res, nil
}
