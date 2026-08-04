package generator

import (
	"iter"
	"math"
	"sync"

	"github.com/DDmit04/test-legion/src/internal/models/exceptions"
)

type PortsRange struct {
	BasePointerGenerator
	start int
	end   int
}

func NewPortsRange(start int, end int) *PortsRange {
	return &PortsRange{
		start: start,
		end:   end,
		BasePointerGenerator: BasePointerGenerator{
			pointer: start,
			start:   start,
			lock:    sync.Mutex{},
		},
	}
}

func (p *PortsRange) Len() int {
	return int(math.Abs(float64(p.end)-float64(p.start)) + 1)
}

func (p *PortsRange) Next() (int, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.pointer > p.end {
		return 0, exceptions.PortsRangeEnds()
	}
	res := p.pointer
	p.pointer++
	return res, nil
}

func (p *PortsRange) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		val, err := p.Next()
		for err == nil {
			if !yield(val) {
				return
			}
			val, err = p.Next()
		}
	}
}

type PortsList struct {
	BasePointerGenerator
	ports []int
}

func (p *PortsList) Len() int {
	return len(p.ports)
}

func NewPortsList(ports []int) *PortsList {
	return &PortsList{
		ports: ports,
		BasePointerGenerator: BasePointerGenerator{
			pointer: 0,
			start:   0,
			lock:    sync.Mutex{},
		},
	}
}

func (p *PortsList) Next() (int, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.pointer >= len(p.ports) {
		return 0, exceptions.PortsRangeEnds()
	}
	res := p.ports[p.pointer]
	p.pointer++
	return res, nil
}

func (p *PortsList) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		val, err := p.Next()
		for err == nil {
			if !yield(val) {
				return
			}
			val, err = p.Next()
		}
	}
}
