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

	posterImageURL = "https://images.microcms-assets.io/assets/4c6b65d46b144491a33e13e269f11289/cbf7a1a07c3e42ed941cb15567af2f4e/poster.jpg"
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
	format := "明日%s曜日に出せるのは\n『%s』\nです。"
	
	now := time.Now()
	wd := now.Weekday()
	switch wd {
	case sun:
		return []linebot.SendingMessage{
			linebot.NewTextMessage("明日月曜日に出せるごみはありません。"),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	case mon:
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "火", "燃やすごみ/燃えないごみ/スプレー缶/乾電池")),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	case tue:
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "水", "缶・びん・ペットボトル/小さな金属類")),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	case wed:
		return []linebot.SendingMessage{
			linebot.NewTextMessage("明日木曜日に出せるごみはありません。"),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	case thu:
		isFirstOrThirdFri := false
		val := float64(now.Day() + 1) / 7
		if val <= 1 || val <= 3 {
			isFirstOrThirdFri = true
		}

		garbage :=  "プラスチック製容器包装"
		if isFirstOrThirdFri {
			garbage += "/古紙/古布"
		}
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "金", garbage)),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	case fri:
		return []linebot.SendingMessage{
			linebot.NewTextMessage(fmt.Sprintf(format, "土", "燃やすごみ/燃えないごみ/スプレー缶/乾電池")),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	case sat:
		return []linebot.SendingMessage{
			linebot.NewTextMessage("明日日曜日に出せるごみはありません。"),
			linebot.NewImageMessage(posterImageURL, posterImageURL),
		}
	}
	return nil
}
