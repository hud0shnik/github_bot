package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hud0shnik/github_bot/internal/telegram"
	"github.com/sirupsen/logrus"
)

// infoResponse - структура статистики пользователя
type infoResponse struct {
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

// commitsResponse - структура количества коммитов за определённый день
type commitsResponse struct {
	Success  bool   `json:"success"`
	Error    string `json:"error"`
	Date     string `json:"date"`
	Username string `json:"username"`
	Commits  int    `json:"commits"`
	Color    int    `json:"color"`
}

// repoResponse - структура статистики репозитория
type repoResponse struct {
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

// SendInfo - функция вывода информации о пользователе GitHub
func SendInfo(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/info <b>[id]</b>\n\nПример:\n/info <b>hud0shnik</b>")
		return
	}

	// Отправка запроса GithubStats Api
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/v2/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка статускода
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
	var user = new(infoResponse)
	json.Unmarshal(body, &user)

	// Отправка данных пользователю
	telegram.SendPict(botUrl, chatId, user.Avatar,
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

// SendCommits - функция вывода количества коммитов
func SendCommits(botUrl string, chatId int, username, date string) {

	// Проверка параметра
	if username == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/commits <b>[id]</b> <b>[date]</b>\n\nПример:\n/commits <b>hud0shnik 2023-02-12</b>\n/commits <b>hud0shnik</b>")
		return
	}

	// Отправка запроса GithubStats Api
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/v2/commits?id=" + username + "&date=" + date)

	// Проверка на ошибку
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка статускода
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
	var user = new(commitsResponse)
	json.Unmarshal(body, &user)

	// Если поле пустое, меняет date на "сегодня"
	if date == "" {
		date = "сегодня"
	}

	// Вывод данных пользователю
	switch user.Color {
	case 1:
		telegram.SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> - <b>%d</b>", date, user.Commits))
		telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIYwmG11bAfndI1wciswTEVJUEdgB2jAAI5AAOtZbwUdHz8lasybOojBA")
	case 2:
		telegram.SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> - <b>%d</b>, неплохо!", date, user.Commits))
		telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIXWmGyDE1aVXGUY6lcjKxx9bOn0JA1AAJOAAOtZbwUIWzOXysr2zwjBA")
	case 3:
		telegram.SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> - <b>%d</b>, отлично!!", date, user.Commits))
		telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIYymG11mMdODUQUZGsQO97V9O0ZLJCAAJeAAOtZbwUvL_TIkzK-MsjBA")
	case 4:
		telegram.SendMsg(botUrl, chatId, fmt.Sprintf("Коммитов за <i>%s</i> - <b>%d</b>, прекрасно!!!", date, user.Commits))
		telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIXXGGyDFClr69PKZXJo9dlYMbyilXLAAI1AAOtZbwU9aVxXMUw5eAjBA")
	default:
		telegram.SendMsg(botUrl, chatId, "Коммитов нет...")
		telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIYG2GzRVNm_d_mVDIOaiLXkGukArlTAAJDAAOtZbwU_-iXZG7hfLsjBA")
	}
}

// SendRepo - функция вывода информации о репозитории
func SendRepo(botUrl string, chatId int, username, reponame string) {

	// Проверка параметров
	if username == "" || reponame == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/repo <b>[username]</b> <b>[reponame]</b>\n\nПример:\n/repo <b>hud0shnik GithubStatsAPI</b>")
		return
	}

	// Отправка запроса GithubStats Api
	resp, err := http.Get("https://githubstatsapi.vercel.app/api/v2/repo?type=string&username=" + username + "&reponame=" + reponame)

	// Проверка на ошибку
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка статускода
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Репозиторий не найден")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
	var repo = new(repoResponse)
	json.Unmarshal(body, &repo)

	// Отправка данных пользователю
	telegram.SendMsg(botUrl, chatId, fmt.Sprintf("Информация о <b>%s/%s</b>\n"+
		"Коммитов <b>%s</b>\n"+
		"Веток <b>%s</b>\n"+
		"Тегов <b>%s</b>\n"+
		"Звёзд <b>%s</b>\n"+
		"Следят <b>%s</b>\n"+
		"Форков <b>%s</b>",
		repo.Username, repo.Reponame, repo.Commits, repo.Branches, repo.Tags, repo.Stars, repo.Watching, repo.Forks))

}
