package exceptions

func PortsRangeEnds(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("no more ports available").
		WithMsg(msg...)
}

func UnexpectedRangeFormat(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("unexpected ports range format").
		WithMsg(msg...)
}

func InvalidPortValue(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("invalid port value").
		WithMsg(msg...)
}

func ZeroPortValue(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("zero ports not accepted").
		WithMsg(msg...)
}

func NegativePortValue(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("negative ports not accepted").
		WithMsg(msg...)
}

func TooHighPortValue(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("ports more than 65535 not accepted").
		WithMsg(msg...)
}
