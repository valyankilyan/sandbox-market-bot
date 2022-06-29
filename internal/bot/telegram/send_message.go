package telegram

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (b *TBot) SendMessage(chatID int64, message string) error {
	log.Printf("Message to %v: %v", chatID, message)
	hc := http.Client{Timeout: 10 * time.Second}
	req, err := b.requestURL("SendMessage")
	if err != nil {
		log.Fatal("cant' make requestURL", err)
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
