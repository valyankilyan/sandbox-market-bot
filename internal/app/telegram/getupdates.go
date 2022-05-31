package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"
)

type JsonUpdates struct {
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

type JsonUpdateRequest struct {
	Offset         int64    `json:"offset"`
	Limit          int64    `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}

func (b *Bot) GetUpdates() (body io.ReadCloser, err error) {
	hc := http.Client{Timeout: 10 * time.Second}
	req, err := b.requestURL("GetUpdates")
	fmt.Println(req)
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}
	updateRequest := JsonUpdateRequest{
		Offset:         config.Conf.Telegram.GetUpdates.Offset,
		Limit:          config.Conf.Telegram.GetUpdates.Limit,
		Timeout:        config.Conf.Telegram.GetUpdates.Timeout,
		AllowedUpdates: config.Conf.Telegram.GetUpdates.AllowedUpdates,
	}
	jsonPost, err := json.Marshal(&updateRequest)
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}

	postBody := bytes.NewReader(jsonPost)
	response, err := hc.Post(req, "application/x-www-form-urlencoded", postBody)
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}
	defer response.Body.Close()

	var target JsonUpdates
	json.NewDecoder(response.Body).Decode(&target)

	got, err := json.MarshalIndent(target, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("getupdates %v", err)
	}

	fmt.Println(string(got))
	return response.Body, nil
}

func (b *Bot) GetMessages(updates io.ReadCloser) error {
	var target JsonUpdates
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
