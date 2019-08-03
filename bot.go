package main

import (
	"os"
	"strconv"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"

	"github.com/Valeyard1/marmitaz-telegram-bot/site"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var token = os.Getenv("TELEGRAM_BOT_TOKEN")
var chatID, _ = strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)

func main() {
	// For better logging
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	c := cron.New()
	// Every weekday on hour 7 through 12 (AM) at second 0 of every minute
	c.AddFunc("0 7-12 * * 1-5", func() {
		var msg tgbotapi.MessageConfig
		if site.TemperoDeMaeIsOpen() {
			msg = tgbotapi.NewMessage(chatID, "Open")
		} else {
			msg = tgbotapi.NewMessage(chatID, "Closed")
		}
		bot.Send(msg)
	})
	c.Start()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Infof("[%s] sent a request to %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Sou uma interface para o site de restaurantes marmitaz.pushsistemas.com.br\nPara mais informações digite /help"
			case "help":
				msg.Text = "Digite /status"
			case "status":
				if site.TemperoDeMaeIsOpen() {
					msg.Text = "O restaurante está aberto. Faça seu pedido"
				} else {
					msg.Text = "O restaurante está fechado."
				}
			default:
				msg.Text = "Q?"
			}
			bot.Send(msg)
		}

	}
}
