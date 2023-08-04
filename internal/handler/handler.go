package handler

import (
	"strings"

	"github.com/hud0shnik/github_bot/internal/api"
	"github.com/hud0shnik/github_bot/internal/commands"
	"github.com/hud0shnik/github_bot/internal/telegram"
)

// Функция генерации и отправки ответа
func Respond(botUrl string, update telegram.Update) {

	// Проверка на сообщение
	if update.Message.Text == "" {
		telegram.SendMsg(botUrl, update.Message.Chat.ChatId, "Пока я воспринимаю только текст")
		return
	}

	// Деление сообщения в слайс
	request := append(strings.Split(update.Message.Text, " "), "", "")

	// Обработчик команд
	switch request[0] {
	case "/info":
		api.SendInfo(botUrl, update.Message.Chat.ChatId, request[1])
	case "/commits":
		api.SendCommits(botUrl, update.Message.Chat.ChatId, request[1], request[2])
	case "/repo":
		api.SendRepo(botUrl, update.Message.Chat.ChatId, request[1], request[2])
	case "/start", "/help":
		commands.Help(botUrl, update.Message.Chat.ChatId)
	default:
		telegram.SendMsg(botUrl, update.Message.Chat.ChatId, "OwO")
	}

}
