package generator

import (
	"iter"
	"sync"
)

type ValueGenerator[T any] interface {
	Next() (T, error)
	Reset()
	Len() int
	Iter() iter.Seq[T]
}

type BasePointerGenerator struct {
	pointer int
	start   int
	lock    sync.Mutex
}

func (l *BasePointerGenerator) Reset() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.pointer = l.start
}
