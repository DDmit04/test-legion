package host

import (
	"fmt"
	"regexp"

	"github.com/DDmit04/test-legion/src/internal/models"
	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/models/generator"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
	"github.com/DDmit04/test-legion/src/internal/tools"
	"github.com/DDmit04/test-legion/src/models/input"
)

type ListReader struct {
	BaseHostReader
}

func NewListReader() *ListReader {
	return &ListReader{
		BaseHostReader: BaseHostReader{
			BaseReader: readers.BaseReader{
				ReadType: input.HostsListInputType,
			},
		},
	}
}

func (r *ListReader) Read(data input.Model) (generator.ValueGenerator[models.DomainInfo], error) {
	sources := make([]string, 0)
	if dataString, ok := data.Value().([]string); ok {
		sources = tools.UniqueStringsList(dataString)
	} else {
		msg := fmt.Sprintf("expected []string, got %T", data)
		return nil, exceptions.InputDataMisMatch(msg)
	}
	res := make([]models.DomainInfo, 0)
	for _, source := range sources {
		validateErr := r.validateHost(source)
		if validateErr != nil {
			validateErr = validateErr.WithMsg(fmt.Sprintf("error host - \"%s\"", source))
			return nil, validateErr
		}
		isHost, _ := regexp.MatchString(domainRegex, source)
		newSource := models.DomainInfo{}
		if isHost {
			ips, err := r.extractDomainIps(source)
			if err != nil {
				return nil, err
			}
			for _, ip := range ips {
				newSource = models.DomainInfo{
					Host: source,
					Ip:   ip,
				}
				res = append(res, newSource)
			}
		} else {
			ip, err := r.validateIp(source)
			if err != nil {
				return nil, err
			}
			newSource = models.DomainInfo{
				Host: "",
				Ip:   ip,
			}
			res = append(res, newSource)
		}

	}
	return generator.NewHostsList(res), nil
}
