package useatomic

import (
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) Incr() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) Get() int64 {
	return c.count
}
