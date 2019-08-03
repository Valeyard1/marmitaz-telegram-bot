package main

import (
	"log"
	"os"

	"github.com/Valeyard1/marmitaz-telegram-bot/site"
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
	if chatID == "" || token == "" {
		log.Fatal("chatID or token not passed")
	}

	log.Println(chatID, token)

	//c := cron.New()
	//c.AddFunc("30 * * * * *", func() {
	//	if site.TemperoDeMaeIsOpen() {
	//		client.SendMessage(chatID, "Open")
	//	} else {
	//		client.SendMessage(chatID, "Closed")
	//	}
	//})
	//c.Start()

	bot.HandleMessage("/start", startHandler)
	bot.HandleMessage("/status", statusHandler)
	bot.HandleMessage("/help", helpHandler)
	err := bot.Start()
	log.Println(err)
}

// StatusHandler handle /status incoming messages
func statusHandler(message *tbot.Message) {
	if site.TemperoDeMaeIsOpen() {
		client.SendMessage(message.Chat.ID, "O restaurante está aberto. Faça seu pedido")
	} else {
		client.SendMessage(message.Chat.ID, "O restaurante está fechado.")
	}
}

// HelpHandler handle /help incoming messages
func helpHandler(message *tbot.Message) {
	client.SendMessage(message.Chat.ID, "Digite /status")
}

// StartHandler handle /start incoming messages
func startHandler(message *tbot.Message) {
	client.SendMessage(message.Chat.ID, "Sou uma interface para o site de restaurantes marmitaz.pushsistemas.com.br\nPara mais informações digite /help")
}
