package src

import (
	"context"
	"fmt"
	"sync"
	"time"

	intModels "github.com/DDmit04/test-legion/src/internal/models"
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/internal/service"
	intProvider "github.com/DDmit04/test-legion/src/internal/services/provider"
	"github.com/DDmit04/test-legion/src/models"
	"github.com/DDmit04/test-legion/src/models/input"
)

type Scanner struct {
	connections int
	timeout     time.Duration
	data        chan models.PortScanResult
	requester   *service.Requester
}

func NewScanner(connections int, timeout time.Duration) *Scanner {
	return &Scanner{
		requester:   service.NewRequester(timeout),
		connections: connections,
		timeout:     timeout,
	}
}

func (s *Scanner) Scan(
	ctx context.Context,
	hostsGen generator.ValueGenerator[intModels.DomainInfo],
	portsGen generator.ValueGenerator[int],
) (chan models.PortScanResult, error) {

	totalChecks := hostsGen.Len() * portsGen.Len()
	fmt.Printf("total checks: %d\n", totalChecks)
	dataChan := make(chan models.PortScanResult, totalChecks)
	sem := make(chan struct{}, s.connections)
	wg := sync.WaitGroup{}

hostLoop:
	for hst := range hostsGen.Iter() {
		for prt := range portsGen.Iter() {
			select {
			case <-ctx.Done():
				break hostLoop
			default:
				wg.Add(1)
				go s.requester.ScanPort(ctx, &wg, sem, dataChan, hst, prt)
			}

		}
		portsGen.Reset()
	}

	go func() {
		wg.Wait()
		close(dataChan)
		close(sem)
	}()

	return dataChan, nil
}

func Scan(ctx context.Context, hosts input.Model, ports input.Model, connections int, timeout time.Duration) (chan models.PortScanResult, error) {
	hostsProvider := intProvider.NewHostsProvider()
	hostsGen, err := hostsProvider.ReadData(hosts)
	if err != nil {
		return nil, err
	}
	portsProvider := intProvider.NewPortsProvider()
	portsGen, err := portsProvider.ReadData(ports)
	if err != nil {
		return nil, err
	}

	scn := NewScanner(connections, timeout)

	channel, err := scn.Scan(ctx, hostsGen, portsGen)

	if err != nil {
		return nil, err
	}

	return channel, nil
}
