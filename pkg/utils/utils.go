package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
	tb "gopkg.in/telebot.v3"
	"pr0.bot/pkg/telegram"
)

const baseUrl = "pr0gramm.com/api/items/get"

// build urls from given category parameters in yaml
func BuildURL(flags int, tags string, promoted int) string {
	uri := &fasthttp.URI{
		DisablePathNormalizing: false,
	}
	uri.SetScheme("https")
	uri.SetHost(baseUrl)
	uri.QueryArgs().Add("flags", fmt.Sprint(flags))
	uri.QueryArgs().Add("promoted", fmt.Sprint(promoted))
	uri.QueryArgs().Add("tags", tags)

	return uri.String()
}

func SendPost(url string, caption string, isAudio bool) ([]tb.Message, error) {
	var media tb.Album

	if isAudio && strings.Contains(url, "mp4") {
		media = tb.Album{
			&tb.Video{
				File:    tb.FromURL(url),
				Caption: fmt.Sprintf("#%s", caption),
			},
		}
	} else {
		media = tb.Album{
			&tb.Photo{
				File:    tb.FromURL(url),
				Caption: fmt.Sprintf("#%s", caption),
			},
		}
	}

	return telegram.TelegramBot.Bot.SendAlbum(telegram.TelegramBot.M.Chat, media)

}

func BuildMediaURL(media string, isAudio bool) string {
	if isAudio && strings.HasSuffix(media, ".mp4") {
		return fmt.Sprintf("https://vid.pr0gramm.com/%s", media)
	}
	return fmt.Sprintf("https://img.pr0gramm.com/%s", media)
}

// generate random waittime to prevent telegram flood
func GenRandomTime() int {
	rand.Seed(time.Now().Unix())
	min := 10
	max := 15

	return rand.Intn(max-min) + min
}

func GenCurrentTime() int64 {
	return time.Now().UnixMilli() / 1000
}
