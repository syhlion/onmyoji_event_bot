package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tucnak/telebot"
)

var (
	db       *sql.DB
	token    = flag.String("t", "", "input telegram bot token")
	location *time.Location
)

func RegisterCommand(m telebot.Message) (msg string, err error) {
	defer func() {
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				msg = "\n這是陰陽師事件機器人，您已註冊過，事件都會通知喔。"
			} else {
				msg = "\n這是陰陽師事件機器人，目前發生錯誤，請再註冊一次"
			}
		}
	}()
	cmd := "INSERT INTO onmyoji (uid,date) VALUES (?,?)"
	tx, err := db.Begin()
	if err != nil {
		return
	}
	stmt, err := tx.Prepare(cmd)
	if err != nil {
		return
	}
	date := time.Now().Format("2006/01/02 15:04:05")

	_, err = stmt.Exec(m.Sender.ID, date)
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
	msg = "\n這是陰陽師事件機器人，恭喜你註冊成功。\n之後 鬥技 妖怪退治 鬼王來襲 陰界之門... 事件都會通知"
	return

}
func UnregisterCommand(m telebot.Message) (msg string, err error) {
	defer func() {
		if err != nil {
			msg = "\n這是陰陽師事件機器人，目前發生錯誤，請聯絡工程師處理"
		}
	}()
	cmd := "DELETE FROM onmyoji WHERE uid = ?"
	_, err = db.Exec(cmd, m.Sender.ID)
	if err != nil {
		return
	}
	msg = "\n 這是陰陽師事件機器人，您已取消註冊。"
	return
}

func HelpCommand(m telebot.Message) (msg string, err error) {
	msg = "\n 這是陰陽師事件機器人，您可以您註冊 \n /register - 註冊\n /unregister - 取消註冊"
	return
}

func init() {
	d, err := sql.Open("sqlite3", "./onmyoji.sqlite3")
	if err != nil {
		return
	}

	sqlStmt := `
	create table if not exists onmyoji (uid PRIMARY KEY,date)
	`
	_, err = d.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	db = d
	location, _ = time.LoadLocation("Asia/Taipei")
}

//鬥技
func Event1(bot *telebot.Bot) {
	t := time.Now()
	t.In(location)
	msg := fmt.Sprintf("現在台灣時間: %s \n鬥技 開始", t.In(location).Format("2006-01-02 15:04:05"))
	query(bot, msg)
}

//妖怪退治
func Event2(bot *telebot.Bot) {
	t := time.Now()
	t.In(location)
	msg := fmt.Sprintf("現在台灣時間: %s \n妖怪退治 開始", t.In(location).Format("2006-01-02 15:04:05"))
	query(bot, msg)
}

//鬼王來襲
func Event3(bot *telebot.Bot) {
	t := time.Now()
	t.In(location)
	msg := fmt.Sprintf("現在台灣時間: %s \n鬼王來襲 開始", t.In(location).Format("2006-01-02 15:04:05"))
	query(bot, msg)
}

//陰界之門
func Event4(bot *telebot.Bot) {
	t := time.Now()
	msg := fmt.Sprintf("現在台灣時間: %s \n陰界之門 開始", t.In(location).Format("2006-01-02 15:04:05"))
	query(bot, msg)
}

//協同鬥技
func Event5(bot *telebot.Bot) {
	t := time.Now()
	msg := fmt.Sprintf("現在台灣時間: %s \n協同鬥技 開始", t.In(location).Format("2006-01-02 15:04:05"))
	query(bot, msg)
}

func query(bot *telebot.Bot, msg string) {
	cmd := "SELECT uid FROM onmyoji"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			continue
		}
		user := telebot.User{ID: id}
		bot.SendMessage(user, msg, nil)

	}
}
func main() {
	flag.Parse()
	bot, err := telebot.NewBot(*token)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Bot Start")
	gocron.Every(1).Day().At("12:00").Do(Event1, bot)
	gocron.Every(1).Day().At("21:00").Do(Event1, bot)
	gocron.Every(1).Day().At("14:00").Do(Event2, bot)
	gocron.Every(1).Monday().At("19:00").Do(Event3, bot)
	gocron.Every(1).Tuesday().At("19:00").Do(Event3, bot)
	gocron.Every(1).Thursday().At("19:00").Do(Event3, bot)
	gocron.Every(1).Wednesday().At("19:00").Do(Event3, bot)
	gocron.Every(1).Friday().At("19:00").Do(Event4, bot)
	gocron.Every(1).Saturday().At("19:00").Do(Event4, bot)
	gocron.Every(1).Sunday().At("19:00").Do(Event4, bot)
	gocron.Every(1).Saturday().At("14:00").Do(Event5, bot)
	gocron.Every(1).Sunday().At("14:00").Do(Event5, bot)
	go func() {
		<-gocron.Start()
	}()

	messages := make(chan telebot.Message)
	bot.Listen(messages, 60*time.Second)
	for message := range messages {
		log.Println(message.Sender.ID, message.Sender.Username, message.Text)
		switch message.Text {
		case "/help":
			s, _ := HelpCommand(message)
			bot.SendMessage(message.Chat, s, nil)
			break
		case "/register":
			s, _ := RegisterCommand(message)
			bot.SendMessage(message.Chat, s, nil)

			break
		case "/unregister":
			s, _ := UnregisterCommand(message)
			bot.SendMessage(message.Chat, s, nil)
			break
		default:
			bot.SendMessage(message.Chat, "請輸入正確指令", nil)
			break
		}

	}
}
