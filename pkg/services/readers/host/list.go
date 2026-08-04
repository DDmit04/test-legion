package host

import (
	"fmt"
	"reflect"
	"regexp"
	"test-legion/pkg/models"
	"test-legion/pkg/models/generator"
	"test-legion/pkg/services/readers"
	"test-legion/tools"
)

type ListReader struct {
	BaseHostReader
}

func NewListReader() *ListReader {
	return &ListReader{
		BaseHostReader: BaseHostReader{
			BaseReader: readers.BaseReader{
				ReadType: reflect.TypeOf(make([]string, 0)),
			},
		},
	}
}

func (r *ListReader) Read(data any) (generator.ValueGenerator[models.DomainInfo], error) {
	sources := tools.UniqueStringsList(data.([]string))
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
