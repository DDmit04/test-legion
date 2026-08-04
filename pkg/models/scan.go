package models

import (
	"net"
	"time"
)

type PortState string

const (
	Open        PortState = "open"
	Closed      PortState = "closed"
	Timeout     PortState = "timeout"
	Unreachable PortState = "unreachable"
	Canceled    PortState = "canceled"
	Error       PortState = "error"
)

type PortScanResult struct {
	Host     string
	IP       net.IP
	Port     int
	State    PortState
	Duration time.Duration
	Err      error
}
