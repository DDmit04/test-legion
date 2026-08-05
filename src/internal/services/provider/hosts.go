package provider

import (
	"fmt"

	"github.com/DDmit04/test-legion/src/internal/models"
	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
	"github.com/DDmit04/test-legion/src/internal/services/readers/host"
	"github.com/DDmit04/test-legion/src/models/input"
)

type HostsProvider struct {
	Readers []readers.Reader[models.DomainInfo]
}

func NewHostsProvider() *HostsProvider {
	return &HostsProvider{
		Readers: []readers.Reader[models.DomainInfo]{
			host.NewListReader(),
		},
	}
}

func (r *HostsProvider) ReadData(hosts input.Model) (generator.ValueGenerator[models.DomainInfo], error) {
	for _, reader := range r.Readers {
		canRead := reader.CanRead(hosts.InputType())
		if canRead {
			return reader.Read(hosts)
		}
	}
	msg := fmt.Sprintf("hosts reader with name: \"%s\" not exists", hosts.InputType())
	return nil, exceptions.ReaderNotFround(msg)
}
