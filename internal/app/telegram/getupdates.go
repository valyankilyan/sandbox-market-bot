package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type JsonGetUpdates struct {
	Status string `json:"status"`
	Result []struct {
		UpdateID int `json:"update_id"`
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

func (b *Bot) GetUpdates() (body io.ReadCloser, err error) {
	hc := http.Client{Timeout: 10 * time.Second}
	req, err := b.requestURL("GetUpdates")
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}
	response, err := hc.Get(req)
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}
	defer response.Body.Close()

	var target JsonGetUpdates
	json.NewDecoder(response.Body).Decode(&target)

	got, err := json.MarshalIndent(target, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}

	fmt.Println(string(got))
	return response.Body, nil
}

func (b *Bot) GetMessages(updates io.ReadCloser) error {
	var target JsonGetUpdates
	json.NewDecoder(updates).Decode(&target)

	messages := make([]Message, 0)
	for _, r := range target.Result {
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
		fmt.Println("message", m.MessageID)
		fmt.Println(m)
	}

	return nil
}
