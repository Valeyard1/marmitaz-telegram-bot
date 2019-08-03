package telegram

import (
	"github.com/Valeyard1/marmitaz-telegram-bot/site"

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

// StartHandler handle /start incoming messages
func StartHandler(message *tbot.Message) {
	client.SendMessage(message.Chat.ID, "Sou uma interface para o site de restaurantes marmitaz.pushsistemas.com.br\nPara mais informações digite /help")
}
