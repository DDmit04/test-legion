package input

import "reflect"

const (
	HostsListInputType Type = "hosts_array"
)

type HostsListInput struct {
	BaseInputModel
}

func CreateHostsListInput(hosts []string) (*HostsListInput, error) {
	res := &HostsListInput{
		BaseInputModel: BaseInputModel{
			value:        hosts,
			inputType:    HostsListInputType,
			expectedType: reflect.TypeOf(hosts),
		},
	}
	err := res.CheckType()
	if err != nil {
		return nil, err
	}
	return res, nil
}
