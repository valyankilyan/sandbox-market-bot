package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

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
