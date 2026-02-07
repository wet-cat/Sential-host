package nfqueue

import (
	"errors"

	nfq "github.com/AkihiroSuda/go-netfilter-queue"
	"sentinel/internal/engine"
	"sentinel/internal/logger"
)

func Run(safe bool) error {
	q := nfq.NewNFQueue(0, 200, nfq.NF_DEFAULT_PACKET_SIZE)
	if q == nil {
		return errors.New("nfqueue init failed (are you root?)")
	}
	defer q.Close()

	logger.Info("nfqueue active")

	for pkt := range q.GetPackets() {
		verdict := engine.AnalyzePacket(pkt.Data)

		// SAFE MODE: never drop packets
		if safe {
			pkt.SetVerdict(nfq.NF_ACCEPT)
		} else {
			if verdict == engine.DROP {
				pkt.SetVerdict(nfq.NF_DROP)
			} else {
				pkt.SetVerdict(nfq.NF_ACCEPT)
			}
		}
	}

	return nil
}

