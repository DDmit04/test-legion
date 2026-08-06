package exceptions

import (
	"errors"
	"strings"
)

type ModuleException struct {
	baseMsg string
	msg     string
	cause   error
}

func NewBaseException() *ModuleException {
	return &ModuleException{
		baseMsg: "Something go wrong",
	}
}

func (e ModuleException) Error() string {
	message := e.baseMsg
	if e.msg != "" {
		message = message + ": " + e.msg
	}
	return message
}

func (e ModuleException) Unwrap() error {
	return e.cause
}

func (e ModuleException) WithBaseMsg(msg string) *ModuleException {
	e.baseMsg = msg
	return &e
}

func (e ModuleException) WithMsg(msg ...string) *ModuleException {
	e.msg = strings.Join(msg, " ")
	return &e
}

func (e ModuleException) WithCause(err error) *ModuleException {
	e.cause = err
	return &e
}

func (e ModuleException) WithUnwrappedCause(err error) *ModuleException {
	var baseEx ModuleException
	if errors.As(err, &baseEx) {
		return &baseEx
	}
	e.cause = err
	return &e
}

func (e ModuleException) WithCauseMsg() *ModuleException {
	if e.cause != nil {
		if e.msg == "" {
			if baseEx := (*ModuleException)(nil); errors.As(e.cause, &baseEx) && baseEx.msg != "" {
				e.msg = baseEx.msg
			} else if e.cause.Error() != "" {
				e.msg = e.cause.Error()
			}
		}
	}
	return &e
}

func UnexpectedError(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("unexpected error").
		WithMsg(msg...)
}
