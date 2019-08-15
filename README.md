# Marmitaz Telegram Bot
**NOTE**: This bot aims to be a brazilian Telegram bot
This is a bot for telegram to interact with the site http://marmitaz.pushsistemas.com.br

It notifies the subscribed users whenever the restaurant opens, which is every weekday around 7~8 AM.  


**TODO**:
* Also notify the menu of the day (If you think it's a good feature, please vote +1 on Valeyard1/marmitaz-telegram-bot#2)

---
For more information, [subscribe](https://t.me/marmitaz_bot) to the bot.


### Contributing
This code is made with go modules, so all the dependencies are stored in the `go.*` files.
To download them, just run `go build`.

With just `go build` the code will not run, you have to specify the TOKEN for your bot, this is made through the `TELEGRAM_BOT_TOKEN` environment
variable.
The following steps can be done to get this working:
```bash
export TELEGRAM_BOT_TOKEN="<token>"
go run bot.go database.go message.go
```