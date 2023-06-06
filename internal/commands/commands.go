package commands

import "github.com/hud0shnik/github_bot/internal/send"

// Функция вывода списка всех команд
func Help(botUrl string, chatId int) {
	send.SendMsg(botUrl, chatId, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/commits <u>username</u> <u>date</u> - коммиты за день\n"+
		"/repo <u>username</u> <u>reponame</u> - статистика репозитория\n"+
		"/info <u>username</u> - информация о пользователе\n")
}
