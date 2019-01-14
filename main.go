package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"
)

func mian() {
	//connect to the bot with token
	bot, err := tgbotapi.NewBotAPI("TOKEH")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// intitialized channel with api
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Fatal(err)
	}
	//read from the channel
	for update := range updates {
		select {
		case update = <-bot.Updates:
			//Information from user who writes to the bot
			UserName := update.Message.From.UserName

			//Chat/dialog ID
			//Can be identificator of user's chat(UserID) and public chat/channel
			ChatID := update.Message.Chat.ID

			//Message text
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			//Answer to the user using his Message
			reply := Text
			//Creating Message
			msg := tgbotapi.NewMessage(ChatID, reply)
			//and send
			bot.SendMessage(msg)

		}
	}
}
