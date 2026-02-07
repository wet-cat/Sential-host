package telemetry

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func Watch(paths []string) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	for _, p := range paths {
		w.Add(p)
	}

	go func() {
		for ev := range w.Events {
			log.Println("[fs]", ev)
		}
	}()

	return nil
}
