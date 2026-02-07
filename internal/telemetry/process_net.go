package telemetry

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Best-effort process lookup via /proc
// Slow, but acceptable for laptops.
func FindProcessByPort(portHex string) (string, error) {
	fds, err := filepath.Glob("/proc/[0-9]*/fd/*")
	if err != nil {
		return "", err
	}

	for _, fd := range fds {
		link, err := ioutil.Readlink(fd)
		if err != nil {
			continue
		}

		if strings.Contains(link, ":"+portHex) {
			parts := strings.Split(fd, "/")
			pid := parts[2]

			cmd, _ := ioutil.ReadFile("/proc/" + pid + "/cmdline")
			return fmt.Sprintf("pid=%s cmd=%s", pid, cmd), nil
		}
	}

	return "", fmt.Errorf("process not found")
}
