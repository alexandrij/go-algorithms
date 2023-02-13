package mutex

import "sync"

type RWMutex struct {
}

func (rw *RWMutex) RLock()
func (rw *RWMutex) RUnlock()

type Counters struct {
	mx sync.RWMutex
	m  map[string]int
}

func (c *Counters) Load(key string) (int, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}
