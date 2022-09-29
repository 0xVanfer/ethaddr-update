package telegram

import (
	"encoding/json"
	"errors"
	"ethaddr-update/config"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/imroc/req"
)

type RobotAdapter struct {
	TeleBot *tgbotapi.BotAPI
}

var AlarmSender *RobotAdapter

func InitTgBot() (err error) {
	cnf := config.GetConfig()
	botToken := cnf.TgToken
	AlarmSender, err = NewBot(botToken)
	return
}

func GetAlarmSender() *RobotAdapter {
	return AlarmSender
}

func Send(botToken string, chatId string, text string) error {
	url := `https://api.telegram.org/bot` + botToken + `/sendMessage?chat_id=` + chatId + `&text=` + text
	_, err := req.Get(url)
	return err
}

func SendComplex(botToken string, chatId int64, text ...any) error {
	alarmBot, err := NewBot(botToken)
	if err != nil {
		return err
	}
	alarmBot.Send(chatId, text...)
	return nil
}

func NewBot(token string) (*RobotAdapter, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("init robot fail:", err)
		return nil, errors.New("")
	}
	TelegramBot := RobotAdapter{TeleBot: bot}
	return &TelegramBot, nil
}

// send telegram
func (r *RobotAdapter) Send(chat_id int64, v ...any) error {
	if r.TeleBot == nil {
		return errors.New("bot.TeleBot is nil")
	}
	msg := tgbotapi.NewMessage(chat_id, fmt.Sprintln(v...))
	msg.ParseMode = "markdown"
	_, err := r.TeleBot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// Read msgs sent to the bot.
//
// Msgs sent by the bot and msgs from group will not be included.
func ReadTgbotPrivateMsgs(botToken string) (msgs []TelegramPrivateMessageStruct, err error) {
	// telegram bot id
	url := "https://api.telegram.org/bot" + botToken + "/getUpdates"
	var v TelegramMessageReq
	// sometimes request will break
	for i := 0; ; i++ {
		// already has content, v is not empty
		if len(v.Result) != 0 {
			break
		}
		// tried too many times
		if i == 3 {
			err = errors.New("request error")
			return
		}
		var r *req.Resp
		r, err = req.Get(url)
		if err != nil {
			continue
		}
		err = r.ToJSON(&v)
		if err != nil {
			continue
		}
	}

	for _, content := range v.Result {
		if content["message"] == "" {
			continue
		}
		bb, _ := json.Marshal(content["message"])
		var msg TelegramPrivateMessageStruct
		err = json.Unmarshal(bb, &msg)
		if err != nil {
			return
		}
		msgs = append(msgs, msg)
	}
	return
}

type TelegramMessageReq struct {
	Ok     bool             `json:"ok"`
	Result []map[string]any `json:"result"`
}

type TelegramPrivateMessageStruct struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date     int    `json:"date"`
	Text     string `json:"text"`
	Entities []struct {
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		Type   string `json:"type"`
	} `json:"entities"`
}
