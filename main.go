package main

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
	config "pr0.bot/pkg/configs"
	_ "pr0.bot/pkg/database"
	"pr0.bot/pkg/logger"

	"pr0.bot/internal/api"
)

func main() {
	logger.SugarLogger.Info("Starting Bot...")
	errGrp, _ := errgroup.WithContext(context.Background())

	for {
		for _, x := range config.Config.Tags.Tags {
			errGrp.Go(func() error { return api.Watcher(x.Flags, x.Tags) })
			time.Sleep(time.Millisecond * 200)
		}

		err := errGrp.Wait()
		if err != nil {
			logger.SugarLogger.Errorln(err)
		}

		time.Sleep(time.Minute * 3)
	}
}
