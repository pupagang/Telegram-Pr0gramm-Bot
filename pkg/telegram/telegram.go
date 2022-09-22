package telegram

import (
	"log"
	"time"

	tb "gopkg.in/telebot.v3"
	"pr0.bot/pkg/configs"
)

type Telegram struct {
	Bot *tb.Bot
	M   tb.Message
}

var TelegramBot Telegram

func init() {
	var err error
	config := tb.Settings{
		Token:  configs.Config.Items.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
	TelegramBot.Bot, err = tb.NewBot(
		config,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var chatRecipient tb.Chat
	chatRecipient.ID = configs.Config.Items.ChannelID
	TelegramBot.M.Chat = &chatRecipient
}
