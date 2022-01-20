package main

import (
	"time"

	config "pr0.bot/pkg/configs"
	_ "pr0.bot/pkg/database"

	"pr0.bot/internal/api"
)

func main() {
	for {
		for _, x := range config.Config.Tags.Tags {
			go api.Watcher(x.Flags, x.Tags, x.Promoted)
			time.Sleep(time.Second * 2)
		}
		time.Sleep(time.Minute * 3)
	}
}
