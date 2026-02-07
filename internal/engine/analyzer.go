package engine

import (
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const maxScore = 10

func AnalyzePacket(raw []byte) uint32 {
	packet := gopacket.NewPacket(raw, layers.LayerTypeIPv4, gopacket.Default)
	if packet == nil {
		return ACCEPT
	}

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return ACCEPT
	}

	ip := ipLayer.(*layers.IPv4)
	ctx := getHost(ip.SrcIP.String())

	if app := packet.ApplicationLayer(); app != nil {
		payload := strings.ToLower(string(app.Payload()))
		ctx.Score += scorePayload(payload)
	}

	ctx.Score += scoreNetworkBehavior(packet)

	if ctx.Score >= maxScore {
		return DROP
	}

	return ACCEPT
}
