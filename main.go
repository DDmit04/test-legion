package main

import (
	"context"
	"fmt"
	"test-legion/pkg/models"
	"test-legion/pkg/services/readers/host"
	"test-legion/pkg/services/readers/port"
	"test-legion/pkg/services/scanner"
	"time"
)

func Scan[T host.ArgsType, V port.ArgTypes](ctx context.Context, hosts T, ports V, connections int, timeout time.Duration) (chan models.PortScanResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

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
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	channel, _ := Scan(ctx, []string{"google.com", "yandex.ru"}, []int{80, 21}, 5, time.Second)
	for val := range channel {
		fmt.Println(val)
	}

}
