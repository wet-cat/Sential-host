package engine

import "strings"

var badPatterns = []string{
	"powershell -enc",
	"/bin/bash -i",
	"nc -e",
	"cmd.exe /c"
}

func scorePayload(payload string) int {
	score := 0

	for _, sig := range badPatterns {
		if strings.Contains(payload, sig) {
			score += 4
		}
	}

	// TODO: base64 decode & inspect
	// TODO: entropy detection

	return score
}
