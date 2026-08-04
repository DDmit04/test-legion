package exceptions

func HostsRangeEnds(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("no more hosts available").
		WithMsg(msg...)
}

func InvalidHost(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("invalid host").
		WithMsg(msg...)
}

func InvalidIp(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("invalid ip").
		WithMsg(msg...)
}

func HostIpsReadError(msg ...string) *ModuleException {
	return NewBaseException().
		WithBaseMsg("host ips read error").
		WithMsg(msg...)
}
