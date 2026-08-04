package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DDmit04/test-legion/pkg/models"
	"github.com/DDmit04/test-legion/pkg/models/input"
	"github.com/DDmit04/test-legion/pkg/services/readers/host"
	"github.com/DDmit04/test-legion/pkg/services/readers/port"
	"github.com/DDmit04/test-legion/pkg/services/scanner"
)

func Scan(ctx context.Context, hosts input.Model, ports input.Model, connections int, timeout time.Duration) (chan models.PortScanResult, error) {
	hostsGen, err := host.ReadHosts(hosts)
	if err != nil {
		return nil, err
	}
	portsGen, err := port.ReadPorts(ports)
	if err != nil {
		return nil, err
	}

	scn := scanner.NewScanner(connections, timeout)

	channel, err := scn.Scan(ctx, hostsGen, portsGen)

	if err != nil {
		return nil, err
	}

	return channel, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hosts, _ := input.CreateHostsListInput([]string{"yandex.ru", "google.com"})
	ports, _ := input.CreatePortsListInput([]int{80, 21})
	channel, _ := Scan(ctx, hosts, ports, 1, 1*time.Second)
	for val := range channel {
		fmt.Println(val)
	}
}
