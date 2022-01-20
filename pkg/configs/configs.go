package configs

import (
	"os"

	"gopkg.in/yaml.v3"
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

	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&Config); err != nil {
		panic(err)
	}
}
