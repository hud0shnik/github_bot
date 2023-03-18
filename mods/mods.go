package mods

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// Структуры для работы с Telegram API

type TelegramResponse struct {
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat    Chat    `json:"chat"`
	Text    string  `json:"text"`
	Sticker Sticker `json:"sticker"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type Sticker struct {
	File_id string `json:"file_id"`
}

// Структуры для работы с GithubStatsApi

type InfoResponse struct {
	Success       bool   `json:"success"`
	Error         string `json:"error"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	Followers     string `json:"followers"`
	Following     string `json:"following"`
	Repositories  string `json:"repositories"`
	Packages      string `json:"packages"`
	Stars         string `json:"stars"`
	Contributions string `json:"contributions"`
	Avatar        string `json:"avatar"`
}

type CommitsResponse struct {
	Error    string `json:"error"`
	Date     string `json:"date"`
	Username string `json:"username"`
	Commits  int    `json:"commits"`
	Color    int    `json:"color"`
}

// Функция вывода информации о пользователе GitHub
func SendInfo(botUrl string, chatId int, username string) {

	// Значение по дефолту
	if username == "" {
		username = "hud0shnik"
	}

	// Отправка запроса
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var user = new(InfoResponse)
	json.Unmarshal(body, &user)

	// Проверка респонса
	if !user.Success {
		SendMsg(botUrl, chatId, user.Error)
		return
	}

	// Отправка данных пользователю
	SendPict(botUrl, chatId, user.Avatar,
		"Информация о "+user.Username+":\n"+
			"Имя "+user.Name+"\n"+
			"Поставленных звезд "+user.Stars+" ⭐\n"+
			"Подписчиков "+user.Followers+" 🤩\n"+
			"Подписок "+user.Following+" 🕵️\n"+
			"Репозиториев "+user.Repositories+" 📘\n"+
			"Пакетов "+user.Packages+" 📦\n"+
			"Контрибуций за год "+user.Contributions+" 🟩\n"+
			"Ссылка на аватар:\n "+user.Avatar)
}

// Функция вывода списка всех команд
func Help(botUrl string, chatId int) {
	SendMsg(botUrl, chatId, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/info <u>username</u> - информация о пользователе\n")
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
