package engine

import (
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	hosts = make(map[string]*HostContext)
)

// TODO: expire old entries
func getHost(ip string) *HostContext {
	mu.Lock()
	defer mu.Unlock()

	h, ok := hosts[ip]
	if !ok {
		h = &HostContext{}
		hosts[ip] = h
	}

	h.LastSeen = time.Now()
	h.RequestCount++

	return h
}
