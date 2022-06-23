package telegram_old

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
)

type JsonUpdates struct {
	Status string `json:"status"`
	Result []struct {
		UpdateID int64 `json:"update_id"`
		Message  struct {
			MessageID int64 `json:"message_id"`
			From      struct {
				ID           int64  `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				Username     string `json:"username"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date int64  `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}

type JsonUpdateRequest struct {
	Offset         int64    `json:"offset"`
	Limit          int64    `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}

func (b *Bot) GetUpdates() (err error) {
	hc := http.Client{Timeout: 10 * time.Second}
	if err != nil {
		return fmt.Errorf("getupdates %v", err)
	}

	for {
		url, err := b.urlUpdate()
		if err != nil {
			return fmt.Errorf("error in getUpdates %v", err)
		}

		// log.Println(url)
		response, err := hc.Get(url)
		if err != nil {
			return fmt.Errorf("getupdates %v", err)
		}

		var updates JsonUpdates
		json.NewDecoder(response.Body).Decode(&updates)
		go b.getMessages(updates)

		if l := len(updates.Result); l != 0 {
			config.Conf.Telegram.GetUpdates.Offset = updates.Result[l-1].UpdateID + 1
		}

		response.Body.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

func (b *Bot) urlUpdate() (url string, err error) {
	gu := config.Conf.Telegram.GetUpdates
	urlquery := make([]string, 1)
	urlquery[0], err = b.requestURL("GetUpdates")
	if err != nil {
		return "", err
	}
	urlquery = append(urlquery, "?")
	urlquery = append(urlquery, fmt.Sprintf("offset=%v&", gu.Offset))
	urlquery = append(urlquery, fmt.Sprintf("limit=%v&", gu.Limit))
	urlquery = append(urlquery, fmt.Sprintf("timeout=%v&", gu.Timeout))
	urlquery = append(urlquery, "allowed_updates=[")
	for au := 0; au < len(gu.AllowedUpdates)-1; au++ {
		urlquery = append(urlquery, fmt.Sprintf("%q,", gu.AllowedUpdates[au]))
	}
	urlquery = append(urlquery, fmt.Sprintf("%q", gu.AllowedUpdates[len(gu.AllowedUpdates)-1]))
	urlquery = append(urlquery, "]")
	url = strings.Join(urlquery, "")
	return url, err
}

func (b *Bot) getMessages(updates JsonUpdates) error {
	messages := make([]Message, 0)
	for _, r := range updates.Result {
		var msg Message
		msg.MessageID = r.Message.MessageID

		msg.Chat.ID = r.Message.Chat.ID
		msg.Chat.FirstName = r.Message.Chat.FirstName
		msg.Chat.Username = r.Message.Chat.Username
		msg.Chat.Type = r.Message.Chat.Type

		msg.From.ID = r.Message.From.ID
		msg.From.IsBot = r.Message.From.IsBot
		msg.From.FirstName = r.Message.From.FirstName
		msg.From.Username = r.Message.From.Username
		msg.From.LanguageCode = r.Message.From.LanguageCode

		msg.Date = r.Message.Date
		msg.Text = r.Message.Text

		messages = append(messages, msg)
	}

	for _, m := range messages {
		fmt.Printf("New message from %s: %s\n", m.From.Username, m.Text)
		go b.HandleMessage(m)
	}

	return nil
}

func (b *Bot) HandleMessage(m Message) {
	fmt.Println(m)
	text := strings.Split(m.Text, " ")
	if len(text) == 0 {
		b.notRecognized(m.Chat.ID)
	}

	switch text[0] {
	case "/start":
		b.sendStart(m.Chat.ID)
		b.createUser(m.From)
	case "/help":
		b.sendHelp(m.Chat.ID)
	case "/tinkoff_token":
		b.tinkoffToken(m, text)
	case "/payin":
		b.payIn(m, text)
	case "/currency":
		b.currency(m, text[1:])
	case "/balance":
		b.balance(m)
	default:
		b.notRecognized(m.Chat.ID)
	}
}
