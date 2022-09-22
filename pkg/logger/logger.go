package logger

import (
	"log"

	"go.uber.org/zap"
)

var SugarLogger *zap.SugaredLogger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}

	defer logger.Sync() // flushes buffer, if any
	SugarLogger = logger.Sugar()

}
