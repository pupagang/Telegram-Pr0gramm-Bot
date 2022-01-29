package main

import (
	"sync"
	"time"

	config "pr0.bot/pkg/configs"
	_ "pr0.bot/pkg/database"

	"pr0.bot/internal/api"
)

func main() {
	var wg sync.WaitGroup
	for {
		for _, x := range config.Config.Tags.Tags {
			wg.Add(1)
			go api.Watcher(x.Flags, x.Tags, x.Promoted, &wg)
			time.Sleep(time.Millisecond * 200)
		}
		wg.Wait()
		time.Sleep(time.Minute * 3)
	}
}
