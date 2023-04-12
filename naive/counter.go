package naive

type Counter struct {
	count int64
}

func (c *Counter) Incr() {
	c.count++
}

func (c *Counter) Get() int64 {
	return c.count
}
