package configs

import (
	"os"

	"gopkg.in/yaml.v3"
	"pr0.bot/pkg/logger"
)

var Config *config

// yaml config struct
type config struct {
	Items   item `yaml:"items"`
	Tags    tags `yaml:"tags"`
	TagList []tag
}

type item struct {
	MongoDBURL string `yaml:"mongodb_url"`
	Cookie     string `yaml:"cookie"`
	BotToken   string `yaml:"bot_token"`
	ChannelID  int64  `yaml:"channel"`
}

type tags struct {
	Tags []tag `yaml:"Tags"`
}

type tag struct {
	Flags    int
	Promoted int
	Tags     string
}

func init() {
	filename := "config.yaml"
	file, _ := os.Open(filename)

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			logger.ErrorLogger.Error(err.Error())
		}
	}(file)

	d := yaml.NewDecoder(file)

	if err := d.Decode(&Config); err != nil {
		panic(err)
	}
}
