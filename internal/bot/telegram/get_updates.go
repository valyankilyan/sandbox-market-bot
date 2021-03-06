package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/bot"
)

func (b *TBot) GetUpdates(msgch chan bot.Message) {
	hc := http.Client{Timeout: 10 * time.Second}

	for {
		url, err := b.urlUpdate()
		if err != nil {
			log.Println("error making urlUpdate", err)
			continue
		}

		response, err := hc.Get(url)
		if err != nil {
			log.Println("error making http request", err)
			continue
		}

		var updates jsonUpdates
		err = json.NewDecoder(response.Body).Decode(&updates)
		if err != nil {
			log.Println("error making http request", err)
			continue
		}

		for _, m := range parseTelegramMsgResp(updates) {
			log.Printf("Message from %v (%v): %v\n", m.From.FirstName, m.From.Username, m.Text)
			msgch <- m
		}

		if l := len(updates.Result); l != 0 {
			config.Telegram.GetUpdates.Offset = updates.Result[l-1].UpdateID + 1
		}

		response.Body.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

func (b *TBot) urlUpdate() (url string, err error) {
	gu := config.Telegram.GetUpdates
	urlquery := make([]string, 1)
	urlquery[0], err = b.requestURL("GetUpdates")
	if err != nil {
		log.Println("error makeing requestURL", err)
	}

	// making query with all parameters given in config
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

func parseTelegramMsgResp(updates jsonUpdates) []bot.Message {
	messages := make([]bot.Message, 0)
	for _, r := range updates.Result {
		var msg bot.Message
		msg.MessageID = r.Message.MessageID

		msg.Chat.ID = r.Message.Chat.ID
		msg.Chat.FirstName = r.Message.Chat.FirstName
		msg.Chat.Username = r.Message.Chat.Username
		msg.Chat.Type = r.Message.Chat.Type

		msg.From.TgId = r.Message.From.ID
		msg.From.IsBot = r.Message.From.IsBot
		msg.From.FirstName = r.Message.From.FirstName
		msg.From.Username = r.Message.From.Username
		msg.From.LanguageCode = r.Message.From.LanguageCode

		msg.Date = r.Message.Date
		msg.Text = r.Message.Text

		messages = append(messages, msg)
	}

	return messages
}

type jsonUpdates struct {
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

type jsonUpdateRequest struct {
	Offset         int64    `json:"offset"`
	Limit          int64    `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}
