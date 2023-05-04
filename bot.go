package aibot

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/g91TeJl/AiBot/pkg/endpoint"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func TelegramBot(api string, apiKey string) {
	fmt.Println("start")
	//Создаем бота
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		panic(err)
	}
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		//Проверяем что от пользователья пришло именно текстовое сообщение
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			switch update.Message.Text {
			case "/start":

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm AiBot. I can generate images")
				bot.Send(msg)

			// case "/number_of_users":

			// 	if os.Getenv("DB_SWITCH") == "on" {

			// 		//Присваиваем количество пользоватьелей использовавших бота в num переменную
			// 		num, err := service.Data.GetNumberOfUsers()
			// 		if err != nil {

			// 			//Отправлем сообщение
			// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error.")
			// 			bot.Send(msg)
			// 		}

			// 		//Создаем строку которая содержит колличество пользователей использовавших бота
			// 		ans := fmt.Sprintf("%d peoples used me for search information in Wikipedia", num)

			// 		//Отправлем сообщение
			// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
			// 		bot.Send(msg)
			// 	} else {

			// 		//Отправлем сообщение
			// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database not connected, so i can't say you how many peoples used me.")
			// 		bot.Send(msg)
			// 	}
			default:
				var count int
				message := update.Message.Text
				fmt.Println(message)
				nc := strings.Split(message, " ")
				if len(nc) != 1 {
					count, _ = strconv.Atoi(nc[1])
				} else {
					count = 1
				}
				wait := tgbotapi.NewMessage(update.Message.Chat.ID, "please Wait")
				bot.Send(wait)
				// в отдельной горутине
				for i := 1; i <= count; i++ {
					endpoint.GenImage(nc[0], i, apiKey)
					l := strconv.Itoa(i)
					data, err := ioutil.ReadFile(fmt.Sprintf("./out/%s_%s.png", nc[0], l))
					if err != nil {
						log.Println(err)
					}
					b := tgbotapi.FileBytes{Name: "image.png", Bytes: data}
					msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, b)
					bot.Send(msg)
				}
			}

		} else {

			//Отправлем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
			bot.Send(msg)
		}
	}
}
