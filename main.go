package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const (
	sun = iota
	mon
	tue
	wed
	thu
	fri
	sat
)

func main() {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(fmt.Errorf("linebot.New(): %w", err))
	}
	
	msgs := newMessages()
	if _, err := bot.BroadcastMessage(msgs...).Do(); err != nil {
		log.Fatal(fmt.Errorf("bot.PushMessage(): %w", err))
	}
	os.Exit(0)
}

func newMessages() (messages []linebot.SendingMessage) {
	format := "明日出せるのは\n『%s』\nです。"
	
	now := time.Now().Add(48*time.Hour)
	wd := now.Weekday()
	switch wd {
	case mon:
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "燃やすごみ/燃えないごみ/スプレー缶/乾電池")),
		}
	case tue:
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "缶・びん・ペットボトル/小さな金属類")),
		}
	case thu:
		isFirstOrThirdFri := false
		val := (now.Day() + 1) / 7
		if val <= 1 || val <= 3 {
			isFirstOrThirdFri = true
		}

		garbage :=  "プラスチック製容器包装"
		if isFirstOrThirdFri {
			garbage += "/古紙/古布"
		}
		return []linebot.SendingMessage{
			linebot.NewTextMessage(garbage),
		}
	case fri:
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "燃やすごみ/燃えないごみ/スプレー缶/乾電池")),
		}
	default:
		return []linebot.SendingMessage{
			linebot.NewTextMessage("明日出せるごみはありません。"),
		}
	}
}
