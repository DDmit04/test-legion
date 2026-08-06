package services

import (
	"context"
	"errors"
	"net"
	"strconv"
	"sync"
	"time"

	internalModels "github.com/DDmit04/test-legion/src/internal/models"
	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
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
	domain internalModels.DomainInfo,
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

	portScanChan := make(chan models.PortScanResult)
	go r.scanPort(ctx, domain, port, portScanChan)

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
	case portScanRes, ok := <-portScanChan:
		if !ok {
			dataChannel <- models.PortScanResult{
				Host:     domain.Host,
				Port:     port,
				IP:       domain.Ip,
				State:    models.Error,
				Duration: 0,
				Err:      exceptions.UnexpectedError("internal data channel closed before get data"),
			}
		}
		dataChannel <- portScanRes
	}
}

func (r *Requester) scanPort(
	ctx context.Context,
	domain internalModels.DomainInfo,
	port int,
	outChan chan models.PortScanResult,
) {

	reqCtx, reqCancel := context.WithTimeout(ctx, r.timeout)
	defer reqCancel()

	target := domain.Host
	if len(target) == 0 {
		target = domain.Ip.String()
	}
	addr := net.JoinHostPort(target, strconv.Itoa(port))
	var d net.Dialer
	start := time.Now()
	conn, err := d.DialContext(reqCtx, "tcp", addr)
	status := r.getPortStatus(err)
	if conn != nil {
		conn.Close()
	}
	duration := time.Since(start)
	outChan <- models.PortScanResult{
		Host:     domain.Host,
		Port:     port,
		IP:       domain.Ip,
		State:    status,
		Duration: duration,
		Err:      err,
	}
	close(outChan)
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
