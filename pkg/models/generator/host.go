package generator

import (
	"iter"
	"sync"

	"github.com/DDmit04/test-legion/pkg/models"
	"github.com/DDmit04/test-legion/pkg/models/exceptions"
)

type HostsList struct {
	BasePointerGenerator
	hosts []models.DomainInfo
}

func NewHostsList(hosts []models.DomainInfo) *HostsList {
	return &HostsList{
		BasePointerGenerator: BasePointerGenerator{
			pointer: 0,
			lock:    sync.Mutex{},
		},
		hosts: hosts,
	}
}

func (l *HostsList) Len() int {
	return len(l.hosts)
}

func (l *HostsList) Next() (models.DomainInfo, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.pointer == len(l.hosts) {
		return models.DomainInfo{}, exceptions.HostsRangeEnds()
	}
	res := l.hosts[l.pointer]
	l.pointer++
	return res, nil
}

func (l *HostsList) Iter() iter.Seq[models.DomainInfo] {
	return func(yield func(models.DomainInfo) bool) {
		val, err := l.Next()
		for err == nil {
			if !yield(val) {
				return
			}
			val, err = l.Next()
		}
	}
}
