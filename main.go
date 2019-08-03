package main

import (
	"log"
	"marmitaz-telegram-bot/site"
	"marmitaz-telegram-bot/telegram"
	"os"

	"github.com/robfig/cron"
	"github.com/yanzay/tbot/v2"
)

var (
	chatID = os.Getenv("TELEGRAM_CHAT_ID")
	token  = os.Getenv("TELEGRAM_BOT_TOKEN")
	client *tbot.Client
)

func main() {
	bot := tbot.New(token, tbot.WithLogger(tbot.BasicLogger{}))
	client = bot.Client()

	c := cron.New()
	c.AddFunc("30 * * * * *", func() {
		if site.TemperoDeMaeIsOpen() {
			client.SendMessage(chatID, "Open")
		} else {
			client.SendMessage(chatID, "Closed")
		}
	})
	c.Start()

	log.Println(c.Entries())

	bot.HandleMessage("/status", telegram.StatusHandler)
	bot.HandleMessage("/help", telegram.HelpHandler)
	bot.Start()
}
