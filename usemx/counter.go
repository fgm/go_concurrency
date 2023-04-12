package usemx

import (
	"sync"
)

type Counter struct {
	mx    sync.Mutex
	count int64
}

func (c *Counter) Incr() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.count++
}

func (c *Counter) Get() int64 {
	return c.count
}
