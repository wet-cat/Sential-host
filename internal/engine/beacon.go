package engine

import "time"

type BeaconTracker struct {
	LastSeen  time.Time
	Intervals []time.Duration
}

// TODO: cap memory usage
func (b *BeaconTracker) Add(t time.Time) {
	if !b.LastSeen.IsZero() {
		b.Intervals = append(b.Intervals, t.Sub(b.LastSeen))
	}
	b.LastSeen = t
}

func (b *BeaconTracker) LooksLikeBeacon() bool {
	if len(b.Intervals) < 5 {
		return false
	}

	first := b.Intervals[0]
	for _, d := range b.Intervals[1:] {
		if abs(d-first) > time.Second {
			return false
		}
	}
	return true
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
