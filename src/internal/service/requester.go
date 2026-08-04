package service

import (
	"context"
	"errors"
	"net"
	"strconv"
	"sync"
	"time"

	models2 "github.com/DDmit04/test-legion/src/internal/models"
	"github.com/DDmit04/test-legion/src/models"
)

type Requester struct {
	timeout time.Duration
}

func NewRequester(timeout time.Duration) *Requester {
	return &Requester{timeout: timeout}
}

func (r *Requester) ScanPort(
	ctx context.Context,
	wg *sync.WaitGroup,
	sem chan struct{},
	dataChannel chan models.PortScanResult,
	domain models2.DomainInfo,
	port int,
) {

	sem <- struct{}{}
	defer func() {
		<-sem
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		return
	default:

	}

	select {
	case <-ctx.Done():
		dataChannel <- models.PortScanResult{
			Host:     domain.Host,
			Port:     port,
			IP:       domain.Ip,
			State:    models.Canceled,
			Duration: 0,
			Err:      nil,
		}
	default:
		start := time.Now()
		target := domain.Host
		if len(target) == 0 {
			target = domain.Ip.String()
		}
		addr := net.JoinHostPort(target, strconv.Itoa(port))

		reqCtx, reqCancel := context.WithTimeout(context.Background(), r.timeout)
		defer reqCancel()

		var d net.Dialer
		conn, err := d.DialContext(reqCtx, "tcp", addr)
		status := r.getPortStatus(err)
		if conn != nil {
			conn.Close()
		}
		duration := time.Since(start)
		select {
		case <-ctx.Done():
			dataChannel <- models.PortScanResult{
				Host:     domain.Host,
				Port:     port,
				IP:       domain.Ip,
				State:    models.Canceled,
				Duration: 0,
				Err:      nil,
			}

			return
		default:
			dataChannel <- models.PortScanResult{
				Host:     domain.Host,
				Port:     port,
				IP:       domain.Ip,
				State:    status,
				Duration: duration,
				Err:      err,
			}
			return
		}
	}
}

func (r *Requester) getPortStatus(err error) models.PortState {
	var netErr net.Error
	if err == nil {
		return models.Open
	}
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return models.Timeout
		}
		if !netErr.Temporary() {
			return models.Closed
		}
		return models.Unreachable
	}
	return models.Error
}
