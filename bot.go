package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/syhlion/gocron"
	"github.com/tucnak/telebot"
)

var (
	db                      *sql.DB
	location                *time.Location
	ONMYOJI_EVENT_BOT_TOKEN string
	dict_tw                 = map[string]string{
		EVENT_S1: "鬥技",
		EVENT_S2: "妖怪退治",
		EVENT_S3: "鬼王來襲",
		EVENT_S4: "陰界之門",
		EVENT_S5: "協同鬥技",
		EVENT_U1: "鬥技",
		EVENT_U2: "妖怪退治",
		EVENT_U3: "鬼王來襲",
		EVENT_U4: "陰界之門",
		EVENT_U5: "協同鬥技",
	}
)

const (
	//鬥技
	EVENT_S1 = "/s1"
	EVENT_U1 = "/u1"
	//妖怪退治
	EVENT_S2 = "/s2"
	EVENT_U2 = "/u2"
	//鬼王來襲
	EVENT_S3 = "/s3"
	EVENT_U3 = "/u3"
	//陰界之門
	EVENT_S4 = "/s4"
	EVENT_U4 = "/u4"
	//協同鬥技
	EVENT_S5 = "/s5"
	EVENT_U5 = "/u5"
)

func RegisterCommand(id int, event string) (msg string) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "UNIQUE") {
				msg = fmt.Sprintf("\n您已註冊 [%s] 過。", dict_tw[event])
			} else {
				msg = "\n這是陰陽師事件機器人，目前發生錯誤，請再註冊一次"
			}
		}
	}()
	cmd := "INSERT INTO onmyoji (date,uid,event_type) VALUES (?,?,?)"
	tx, err := db.Begin()
	if err != nil {
		return
	}
	stmt, err := tx.Prepare(cmd)
	if err != nil {
		return
	}
	date := time.Now().Format("2006/01/02 15:04:05")

	_, err = stmt.Exec(date, id, event)
	if err != nil {
		tx.Rollback()
		stmt.Close()
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		stmt.Close()
		return
	}
	msg = fmt.Sprintf("\n註冊 [%s] 成功。", dict_tw[event])
	return

}
func UnregisterCommand(id int, event string) (msg string) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			msg = "\n目前發生錯誤，請聯絡工程師處理"
		}
	}()
	cmd := "DELETE FROM onmyoji WHERE uid = ? and event_type = ?"
	_, err = db.Exec(cmd, id, event)
	if err != nil {
		log.Println(err)
		return
	}
	msg = fmt.Sprintf("\n取消%s註冊。", dict_tw[event])
	return
}

func ListCommand(id int) (msg string) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			msg = "\n目前發生錯誤，請聯絡工程師處理"
		}
	}()
	cmd := "SELECT event_type FROM onmyoji WHERE uid = ?"
	rows, err := db.Query(cmd, id)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	var buffer bytes.Buffer
	buffer.WriteString("這是您目前訂閱的事件列表\n\n")
	var eventType string
	i := 0
	for rows.Next() {
		i++
		err = rows.Scan(&eventType)
		if err != nil {
			log.Println(err)
			continue
		}
		buffer.WriteString(dict_tw[eventType] + "\n")
	}
	if i != 0 {
		msg = buffer.String()

	} else {
		msg = "您目前沒有訂閱事件"
	}
	return
}

func init() {
	os.Mkdir("db/", 0755)
	d, err := sql.Open("sqlite3", "db/onmyoji.sqlite3")
	if err != nil {
		return
	}

	sqlStmt := `
	create table if not exists onmyoji (date,uid,event_type,UNIQUE(uid,event_type));
	`
	_, err = d.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	db = d
	location, err = time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Println(err)
		return
	}
	ONMYOJI_EVENT_BOT_TOKEN = os.Getenv("ONMYOJI_EVENT_BOT_TOKEN")
}

func callbacks(bot *telebot.Bot) {
	for callback := range bot.Callbacks {
		var cbr *telebot.CallbackResponse
		switch callback.Data {
		case EVENT_S1:
			msg := RegisterCommand(callback.Sender.ID, EVENT_S1)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_S2:
			msg := RegisterCommand(callback.Sender.ID, EVENT_S2)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_S3:
			msg := RegisterCommand(callback.Sender.ID, EVENT_S3)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_S4:
			msg := RegisterCommand(callback.Sender.ID, EVENT_S4)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_S5:
			msg := RegisterCommand(callback.Sender.ID, EVENT_S5)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_U1:
			msg := UnregisterCommand(callback.Sender.ID, EVENT_S1)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_U2:
			msg := UnregisterCommand(callback.Sender.ID, EVENT_S2)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_U3:
			msg := UnregisterCommand(callback.Sender.ID, EVENT_S3)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_U4:
			msg := UnregisterCommand(callback.Sender.ID, EVENT_S4)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		case EVENT_U5:
			msg := UnregisterCommand(callback.Sender.ID, EVENT_S5)
			cbr = &telebot.CallbackResponse{
				CallbackID: callback.ID,
				Text:       msg,
				ShowAlert:  true,
			}
		}
		bot.AnswerCallbackQuery(&callback, cbr)
	}
}

func messages(bot *telebot.Bot) {
	for message := range bot.Messages {
		log.Println(message.ID, message.Sender.ID, message.Sender.Username, message.Text)
		switch message.Text {
		case "/subscribe":
			bot.SendMessage(message.Chat, "請選擇要註冊的事件", &telebot.SendOptions{
				ReplyMarkup: telebot.ReplyMarkup{
					OneTimeKeyboard: true,
					InlineKeyboard: [][]telebot.KeyboardButton{
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_S1],
								Data: EVENT_S1,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_S2],
								Data: EVENT_S2,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_S3],
								Data: EVENT_S3,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_S4],
								Data: EVENT_S4,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_S5],
								Data: EVENT_S5,
							},
						},
					},
				},
			})

			break
		case "/unsubscribe":
			bot.SendMessage(message.Chat, "請選擇要取消的事件", &telebot.SendOptions{
				ReplyMarkup: telebot.ReplyMarkup{
					OneTimeKeyboard: true,
					InlineKeyboard: [][]telebot.KeyboardButton{
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_U1],
								Data: EVENT_U1,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_U2],
								Data: EVENT_U2,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_U3],
								Data: EVENT_U3,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_U4],
								Data: EVENT_U4,
							},
						},
						[]telebot.KeyboardButton{
							telebot.KeyboardButton{
								Text: dict_tw[EVENT_U5],
								Data: EVENT_U5,
							},
						},
					},
				},
			})

			break
		case "/start":
			option := &telebot.SendOptions{
				DisableWebPagePreview: true,
				ReplyMarkup: telebot.ReplyMarkup{
					//OneTimeKeyboard: true,
					CustomKeyboard: [][]string{
						[]string{"/subscribe"},
						[]string{"/unsubscribe"},
						[]string{"/list"},
					},
				},
			}
			text := "歡迎使用陰陽師事件通知機器人\n\n問題回報\nhttps://github.com/syhlion/onmyoji_event_bot/issues\n\n"
			bot.SendMessage(message.Chat, text, option)
		case "/help":
			option := &telebot.SendOptions{
				DisableWebPagePreview: true,
			}
			text := "歡迎使用陰陽師事件通知機器人\n\n\n問題回報\nhttps://github.com/syhlion/onmyoji_event_bot/issues\n\n命令列表\n/subscribe - 註冊遊戲事件\n/unsubscribe - 取消遊戲事件\n/list - 列出目前所有訂閱事件"
			bot.SendMessage(message.Chat, text, option)
		case "/list":
			msg := ListCommand(message.Sender.ID)
			bot.SendMessage(message.Chat, msg, nil)

		default:
			bot.SendMessage(message.Chat, "請輸入正確指令", nil)
			break
		}

	}
}
func Event(bot *telebot.Bot, event string) {
	cmd := "SELECT uid FROM onmyoji WHERE event_type = ?"
	rows, err := db.Query(cmd, event)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	t := time.Now()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			continue
		}

		msg := fmt.Sprintf("現在台灣時間: %s \n%s 開始", t.In(location).Format("2006-01-02 15:04:05"), dict_tw[event])
		user := telebot.User{ID: id}
		bot.SendMessage(user, msg, nil)

	}
}
func main() {
	bot, err := telebot.NewBot(ONMYOJI_EVENT_BOT_TOKEN)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Bot Start")
	gocron.ChangeLoc(location)
	//鬥技
	gocron.Every(1).Days().At("12:00").Do(Event, bot, EVENT_S1)
	gocron.Every(1).Days().At("21:00").Do(Event, bot, EVENT_S1)
	//妖怪退治
	gocron.Every(1).Days().At("13:00").Do(Event, bot, EVENT_S2)
	gocron.Every(1).Days().At("20:00").Do(Event, bot, EVENT_S2)
	//鬼王來襲
	gocron.Every(1).Monday().At("19:00").Do(Event, bot, EVENT_S3)
	gocron.Every(1).Tuesday().At("19:00").Do(Event, bot, EVENT_S3)
	gocron.Every(1).Thursday().At("19:00").Do(Event, bot, EVENT_S3)
	gocron.Every(1).Wednesday().At("19:00").Do(Event, bot, EVENT_S3)
	//陰界之門
	gocron.Every(1).Friday().At("19:00").Do(Event, bot, EVENT_S4)
	gocron.Every(1).Saturday().At("19:00").Do(Event, bot, EVENT_S4)
	gocron.Every(1).Sunday().At("19:00").Do(Event, bot, EVENT_S4)
	//協同鬥技
	gocron.Every(1).Saturday().At("14:00").Do(Event, bot, EVENT_S5)
	gocron.Every(1).Sunday().At("14:00").Do(Event, bot, EVENT_S5)
	go func() {
		<-gocron.Start()
	}()

	bot.Messages = make(chan telebot.Message)
	bot.Queries = make(chan telebot.Query)
	bot.Callbacks = make(chan telebot.Callback)
	//bot.Listen(messages, 60*time.Second)
	go messages(bot)
	go callbacks(bot)
	bot.Start(60 * time.Second)

}
