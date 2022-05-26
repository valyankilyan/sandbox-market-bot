package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JsonGetUpdates struct {
	Status string `json:"status"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
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
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}

func (b *Bot) GetUpdates() error {
	hc := http.Client{Timeout: 10 * time.Second}
	req, err := b.requestURL("GetUpdates")
	if err != nil {
		return fmt.Errorf("getupdates %v", err)
	}
	response, err := hc.Get(req)
	if err != nil {
		return fmt.Errorf("getupdates %v", err)
	}
	defer response.Body.Close()

	var target JsonGetUpdates
	json.NewDecoder(response.Body).Decode(&target)

	got, err := json.MarshalIndent(target, "", "\t")
	if err != nil {
		return fmt.Errorf("getupdates %v", err)
	}

	fmt.Println(string(got))
	return nil
}
