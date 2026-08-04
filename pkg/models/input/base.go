package input

import (
	"fmt"
	"reflect"

	"github.com/DDmit04/test-legion/pkg/models/exceptions"
)

type Type string

type Model interface {
	InputType() Type
	Value() any
	CheckType() error
}

type BaseInputModel struct {
	value        any
	inputType    Type
	expectedType any
}

func (b *BaseInputModel) CheckType() error {
	valueType := reflect.TypeOf(b.value)
	expectedType := reflect.TypeOf(b.expectedType)
	if valueType == expectedType {
		msg := fmt.Sprintf("expected type %v but got %v", expectedType.String(), valueType.String())
		return exceptions.InputDataMisMatch(msg)
	}
	return nil
}

func (b *BaseInputModel) InputType() Type {
	return b.inputType
}

func (b *BaseInputModel) Value() any {
	return b.value
}
