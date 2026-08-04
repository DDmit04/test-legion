package scanner

import (
	"context"
	"sync"
	"test-legion/pkg/models"
	"test-legion/pkg/models/generator"
	"time"
)

type Scanner struct {
	connections int
	timeout     time.Duration
	data        chan models.PortScanResult
	requester   *Requester
}

func NewScanner(connections int, timeout time.Duration) *Scanner {
	return &Scanner{
		requester:   NewRequester(timeout),
		connections: connections,
		timeout:     timeout,
	}
}

func (s *Scanner) Scan(
	ctx context.Context,
	hostsGen generator.ValueGenerator[models.DomainInfo],
	portsGen generator.ValueGenerator[int],
) (chan models.PortScanResult, error) {

	dataChan := make(chan models.PortScanResult, hostsGen.Len()+portsGen.Len())
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
				sem <- struct{}{}
				go s.requester.scanPort(ctx, &wg, sem, dataChan, hst, prt)
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
