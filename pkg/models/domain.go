package models

import "net"

type DomainInfo struct {
	Ip   net.IP
	Host string
}
