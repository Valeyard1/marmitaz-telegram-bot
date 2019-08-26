package main

import (
	"os"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"

	"github.com/Valeyard1/marmitaz-telegram-bot/site"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	openRestaurantKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("marmitaz.com", "https://marmitaz.pushsistemas.com.br/"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Já pedi", "0"),
			tgbotapi.NewInlineKeyboardButtonData("Vou pedir", "1"),
		),
	)
	token         = os.Getenv("TELEGRAM_BOT_TOKEN")
	DATABASE_HOST = os.Getenv("DATABASE_HOST") // File of sqlite to use
)

func main() {
	// For better logging
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	if token == "" {
		log.Fatal("No token has been provided for bot to work. Provide the TELEGRAM_BOT_TOKEN environment variable")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err.Error())
	}

	db := initializeDatabase()
	defer db.Close()

	c := cron.New()
	// Every weekday on hour 7 through 12 (AM) at second 0 of every minute
	err = c.AddFunc("0 * 7-11 * * 1-5", func() {
		var users []User
		db.Find(&users)

		restaurantIsOpen, err := site.TemperoDeMaeIsOpen()
		if err != nil {
			log.Errorf("Failed to get status of restaurant in the cron job\n%v", err.Error())
		}

		for _, subscribed := range users {
			var msg tgbotapi.MessageConfig
			if restaurantIsOpen {
				msg = tgbotapi.NewMessage(subscribed.UserID, "O Restaurante abriu. Faça seu pedido")
				msg.ReplyMarkup = openRestaurantKeyboard
				log.Info("Sent a notification for ", subscribed.Username)
			}
			_, err = bot.Send(msg)
		}
	})
	c.Start()

	log.Infof("Authorized on account %s", bot.Self.UserName)
	log.Info("Listening for new commands")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.CallbackQuery != nil {
			var text string

			switch update.CallbackQuery.Data {
			case "0":
				text = "Até amanhã :D"
				var user User
				db.Model(&user).Where("user_id = ?", update.CallbackQuery.From.ID).Update("order", 1)
			}

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "ok"))

			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text))
		}

		if update.Message == nil {
			continue
		}

		log.Infof("[%s] - (%d) sent a request to %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "subscribe":
				db.Create(&User{Username: update.Message.From.UserName, UserID: update.Message.Chat.ID})
				msg.Text = "Você será notificado assim que o restaurante abrir. Para cancelar digite /cancel"
			case "help":
				msg.Text = helpMessage()
			case "status":
				if isOpen, _ := site.TemperoDeMaeIsOpen(); isOpen == true {
					msg.Text = "O restaurante está aberto. Faça seu pedido."
					msg.ReplyMarkup = openRestaurantKeyboard
				} else {
					msg.Text = "O restaurante está fechado."
				}
			case "cancel":
				db.Where("user_id = ?", update.Message.Chat.ID).Delete(User{})
				db.Unscoped().Delete(&User{})
				msg.Text = "Você não será mais notificado. Para se inscrever novamente digite /subscribe"
			case "querocafe":
				msg.Text = queroCafeMessage()
			default:
				msg.Text = "Opção não disponível, para listar os comandos disponíveis digite /help"
			}
			failedMSG, err := bot.Send(msg)
			if err != nil {
				log.Errorf("Message %s not sent.\n%v", failedMSG, err)
			}
		}

	}
}
