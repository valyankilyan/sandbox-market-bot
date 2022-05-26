package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type Bot struct {
	token string
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func (b *Bot) requestURL(command string) (string, error) {
	if command == "" || b.token == "" {
		return "", fmt.Errorf("something went wrong")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}

func New(token string) *Bot {
	return &Bot{
		token: token,
	}
}

func (b *Bot) SendMessage(chatID int64, message string) error {
	hc := http.Client{Timeout: 10 * time.Second}
	req, err := b.requestURL("SendMessage")
	if err != nil {
		return err
	}

	payload := fmt.Sprintf(`
		{
			"chat_id": %d,
			"text": %q
		}
	`, chatID, message)
	_, err = hc.Post(req, "application/json", bytes.NewBuffer([]byte(payload)))
	return err
}

func (b *Bot) GetUpdates() (*http.Response, error) {
	hc := http.Client{Timeout: 10 * time.Second}
	req, err := b.requestURL("GetUpdates")
	if err != nil {
		return nil, err
	}
	get, err := hc.Get(req)
	buff := make([]byte, 10000)
	body, _ := get.Body.Read(buff)
	fmt.Println("BODY:", body)
	fmt.Println("Buffer:", string(buff))
	return get, err
}
