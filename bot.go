package aibot

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/segmentio/kafka-go"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	broker1 = "localhost:9092"
	//broker2 = "localhost:9094"
)

func produce(ctx context.Context, msg string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1},
		Topic:   "test2",
	})
	key := 0
	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(key)),
			Value: []byte(msg),
		})
		if err != nil {
			log.Fatal(err)
		}

		key++
		//breakew
		return
	}
}

func consume(ctx context.Context) string {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1},
		Topic:   "test2",
		//GroupID: "te2t",
		//StartOffset: kafka.LastOffset,
	})

	r.SetOffset(-1)
	//r.FetchMessage()
	//r.SetOffsetAt(ctx, time.Time{})
	//r.SetOffset(-2)
	fmt.Println(r.Offset())
	for {
		msg, err := r.ReadMessage(ctx)

		if err != nil {
			return err.Error()
		}
		//r.CommitMessages(ctx, msg)
		return string(msg.Value)
	}

}

func TelegramBot(ctx context.Context, api string, apiKey string) {
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
			default:
				//	var count int
				message := update.Message.Text
				go produce(ctx, message)
				//fmt.Println(message)
				// nc := strings.Split(message, " ")
				// if len(nc) != 1 {
				// 	count, _ = strconv.Atoi(nc[1])
				// } else {
				// 	count = 1
				// }
				//time.Sleep(time.Second * 3)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, consume(ctx))
				fmt.Println(msg)
				bot.Send(msg)
				// wait := tgbotapi.NewMessage(update.Message.Chat.ID, "please Wait")
				// bot.Send(wait)
				// // в отдельной горутине
				// for i := 1; i <= count; i++ {
				// 	endpoint.GenImage(nc[0], i, apiKey)
				// 	l := strconv.Itoa(i)
				// 	data, err := ioutil.ReadFile(fmt.Sprintf("./out/%s_%s.png", nc[0], l))
				// 	if err != nil {
				// 		log.Println(err)
				// 	}
				// 	b := tgbotapi.FileBytes{Name: "image.png", Bytes: data}
				// 	msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, b)
				// 	bot.Send(msg)
				// }
			}

		} else {

			//Отправлем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
			bot.Send(msg)
		}
	}
}
