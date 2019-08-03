package telegram

import (
	"marmitaz-telegram-bot/site"

	"github.com/yanzay/tbot/v2"
)

var client *tbot.Client

// StatusHandler handle /status incoming messages
func StatusHandler(message *tbot.Message) {
	if site.TemperoDeMaeIsOpen() {
		client.SendMessage(message.Chat.ID, "O restaurante está aberto. Faça seu pedido")
	} else {
		client.SendMessage(message.Chat.ID, "O restaurante está fechado.")
	}
}

// HelpHandler handle /help incoming messages
func HelpHandler(message *tbot.Message) {
	client.SendMessage(message.Chat.ID, "Digite /status")
}
