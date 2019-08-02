package main

import (
	"marmitaz-telegram-bot/site"
	"os"

	"github.com/yanzay/tbot/v2"
)

var client *tbot.Client

func main() {
	// c := cron.New()
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot := tbot.New(token)
	client = bot.Client()

	bot.HandleMessage("/status", statusHandler)
	bot.HandleMessage("/help", helpHandler)
	bot.Start()
}

func statusHandler(message *tbot.Message) {
	if site.TemperoDeMaeIsOpen() {
		client.SendMessage(message.Chat.ID, "O restaurante está aberto. Faça seu pedido")
	} else {
		client.SendMessage(message.Chat.ID, "O restaurante está fechado.")
	}
}

func helpHandler(message *tbot.Message) {
	client.SendMessage(message.Chat.ID, "Type /status")
}
