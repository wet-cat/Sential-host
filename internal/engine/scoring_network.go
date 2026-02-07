package engine

import "github.com/google/gopacket"

func scoreNetworkBehavior(packet gopacket.Packet) int {
	score := 0

	// TODO: port scan detection
	// TODO: outbound uncommon ports
	// TODO: timing-based heuristics

	return score
}
