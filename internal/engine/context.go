package engine

import "time"

type HostContext struct {
	Score        int
	LastSeen     time.Time
	RequestCount int
}
