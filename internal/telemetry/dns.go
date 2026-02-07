package telemetry

import "unicode"

func LooksRandom(domain string) bool {
	weird := 0
	for _, r := range domain {
		if unicode.IsDigit(r) || r > 127 {
			weird++
		}
	}
	return weird > len(domain)/2
}
