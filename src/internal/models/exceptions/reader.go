package exceptions

func ReaderNotFround(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("reader not found").
		WithMsg(msg...)
}
