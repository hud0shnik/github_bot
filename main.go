package main

import (
	"os"
	"time"

	"github.com/hud0shnik/github_bot/internal/handler"
	"github.com/hud0shnik/github_bot/internal/telegram"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	// Загрузка переменных окружения
	godotenv.Load()

	// Url бота для отправки и приёма сообщений
	botUrl := "https://api.telegram.org/bot" + os.Getenv("TOKEN")
	offSet := 0

	// Уведомление о старте
	logrus.Info("Bot is running")

	// Цикл работы бота
	for {

		// Получение апдейтов
		updates, err := telegram.GetUpdates(botUrl, offSet)
		if err != nil {
			logrus.Fatalf("getUpdates error: %s", err)
			return
		}

		// Обработка апдейтов
		for _, update := range updates {
			handler.Respond(botUrl, update)
			offSet = update.UpdateId + 1
		}

	}
}
