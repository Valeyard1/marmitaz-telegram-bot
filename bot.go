package test

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	msg := tgbotapi.NewMessage(459634780, "APT")
	bot.Send(msg)
	return

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		var msg tgbotapi.MessageConfig
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		log.Println(update.Message.Chat.ID)

		//if site.TemperoDeMaeIsOpen() {
		//	msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//	msg.ReplyToMessageID = update.Message.MessageID
		//} else {
		//	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "TESTEE")
		//	msg.ReplyToMessageID = update.Message.MessageID
		//}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		bot.Send(msg)
	}
}
