package main

// d (map[string]string) [0x11111]

//   cada cliente (goroutine) INCR contador +1 (10 + 11)
//                             | lock | alteraacao | unlock |
//   cada cliente (goroutine) aguardando, aguardando, lock() -- altera -- unlock() 
//                            INCR contador +1  w, w,w,w    | (11 + 12)

import (
	"sync"
)

type CacheD struct {
	d   map[string]string
	sync.RWMutex
}

var gCacheD CacheD

func (c *CacheD) Add(chave, valor string) {
	c.Lock()
	c.d[chave] = valor
	c.Unlock()
}

func (c *CacheD) Del(chave string) {
	c.Lock()
	defer c.Unlock()

	if _, existe := c.d[chave]; existe {
		delete(c.d, chave)
	}
}

func (c *CacheD) Get(chave string) string {
	c.RLock()
	defer c.RUnlock()
	return c.d[chave]
}
