package host

import (
	"fmt"

	"github.com/DDmit04/test-legion/pkg/models"
	"github.com/DDmit04/test-legion/pkg/models/generator"
	"github.com/DDmit04/test-legion/pkg/models/input"
	rdrs "github.com/DDmit04/test-legion/pkg/services/readers"
)

var hostReaders = []rdrs.Reader[models.DomainInfo]{
	NewListReader(),
}

func ReadHosts(hosts input.Model) (generator.ValueGenerator[models.DomainInfo], error) {
	for _, reader := range hostReaders {
		canRead := reader.CanRead(hosts.InputType())
		if canRead {
			return reader.Read(hosts)
		}
	}
	return nil, fmt.Errorf("not expected generator type")
}
