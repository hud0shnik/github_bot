package mods

import (
	"encoding/json"
	"fmt"
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
	Success  bool   `json:"success"`
	Error    string `json:"error"`
	Date     string `json:"date"`
	Username string `json:"username"`
	Commits  int    `json:"commits"`
	Color    int    `json:"color"`
}

type RepoResponse struct {
	Success  bool   `json:"success"`
	Error    string `json:"error"`
	Username string `json:"username"`
	Reponame string `json:"reponame"`
	Commits  string `json:"commits"`
	Branches string `json:"branches"`
	Tags     string `json:"tags"`
	Stars    string `json:"stars"`
	Watching string `json:"watching"`
	Forks    string `json:"forks"`
}

// Функция вывода информации о пользователе GitHub
func SendInfo(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/info <b>[id]</b>\n\nПример:\n/info <b>hud0shnik</b>")
		return
	}

	// Отправка запроса GithubStats Api
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/v2/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var user = new(InfoResponse)
	json.Unmarshal(body, &user)

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса, продолжение выполнения кода
	case 404:
		SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Отправка данных пользователю
	SendPict(botUrl, chatId, user.Avatar,
		"Информация о <b>"+user.Username+"</b>:\n"+
			"Имя "+user.Name+"\n"+
			"Поставленных звезд <b>"+user.Stars+"</b> ⭐\n"+
			"Подписчиков <b>"+user.Followers+"</b> 🤩\n"+
			"Подписок <b>"+user.Following+"</b> 🕵️\n"+
			"Репозиториев <b>"+user.Repositories+"</b> 📘\n"+
			"Пакетов <b>"+user.Packages+"</b> 📦\n"+
			"Контрибуций за год <b>"+user.Contributions+"</b> 🟩\n"+
			"Аватар:\n"+user.Avatar)
}

// Функция вывода количества коммитов
func SendCommits(botUrl string, chatId int, username, date string) {

	// Проверка параметра
	if username == "" {
		SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/commits <b>[id]</b> <b>[date]</b>\n\nПример:\n/commits <b>hud0shnik 2023-02-12</b>\n/commits <b>hud0shnik</b>")
		return
	}

	// Отправка запроса GithubStats Api
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/v2/commits?id=" + username + "&date=" + date)

	// Проверка на ошибку
	if err != nil {
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var user = new(CommitsResponse)
	json.Unmarshal(body, &user)

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса, продолжение выполнения кода
	case 404:
		SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Если поле пустое, меняет date на "сегодня"
	if date == "" {
		date = "сегодня"
	}

	// Вывод данных пользователю
	switch user.Color {
	case 1:
		SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> <b>%d</b>", date, user.Commits))
		SendStck(botUrl, chatId, "CAACAgIAAxkBAAIYwmG11bAfndI1wciswTEVJUEdgB2jAAI5AAOtZbwUdHz8lasybOojBA")
	case 2:
		SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> <b>%d</b>, неплохо!", date, user.Commits))
		SendStck(botUrl, chatId, "CAACAgIAAxkBAAIXWmGyDE1aVXGUY6lcjKxx9bOn0JA1AAJOAAOtZbwUIWzOXysr2zwjBA")
	case 3:
		SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> <b>%d</b>, отлично!!", date, user.Commits))
		SendStck(botUrl, chatId, "CAACAgIAAxkBAAIYymG11mMdODUQUZGsQO97V9O0ZLJCAAJeAAOtZbwUvL_TIkzK-MsjBA")
	case 4:
		SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> <b>%d</b>, прекрасно!!!", date, user.Commits))
		SendStck(botUrl, chatId, "CAACAgIAAxkBAAIXXGGyDFClr69PKZXJo9dlYMbyilXLAAI1AAOtZbwU9aVxXMUw5eAjBA")
	default:
		SendMsg(botUrl, chatId, "Коммитов нет...")
		SendStck(botUrl, chatId, "CAACAgIAAxkBAAIYG2GzRVNm_d_mVDIOaiLXkGukArlTAAJDAAOtZbwU_-iXZG7hfLsjBA")
	}
}

// Функция вывода информации о репозитории
func SendRepo(botUrl string, chatId int, username, reponame string) {

	// Проверка параметров
	if username == "" || reponame == "" {
		SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/repo <b>[username]</b> <b>[reponame]</b>\n\nПример:\n/repo <b>hud0shnik GithubStatsAPI</b>")
		return
	}

	// Отправка запроса GithubStats Api
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/v2/repo?type=string&username=" + username + "&reponame=" + reponame)

	// Проверка на ошибку
	if err != nil {
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var repo = new(RepoResponse)
	json.Unmarshal(body, &repo)

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса, продолжение выполнения кода
	case 404:
		SendMsg(botUrl, chatId, "Репозиторий не найден")
		return
	case 400:
		SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Отправка данных пользователю
	SendMsg(botUrl, chatId, fmt.Sprintf("Информация о <b>%s/%s</b>\n"+
		"Коммитов <b>%s</b>\n"+
		"Веток <b>%s</b>\n"+
		"Тегов <b>%s</b>\n"+
		"Звёзд <b>%s</b>\n"+
		"Следят <b>%s</b>\n"+
		"Форков <b>%s</b>",
		repo.Username, repo.Reponame, repo.Commits, repo.Branches, repo.Tags, repo.Stars, repo.Watching, repo.Forks))

}

// Функция вывода списка всех команд
func Help(botUrl string, chatId int) {
	SendMsg(botUrl, chatId, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/commits <u>username</u> <u>date</u> - коммиты за день\n"+
		"/repo <u>username</u> <u>reponame</u> - статистика репозитория\n"+
		"/info <u>username</u> - информация о пользователе\n")
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
