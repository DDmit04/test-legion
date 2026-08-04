package exceptions

func InputDataMisMatch(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("invalid input").
		WithMsg(msg...)
}
