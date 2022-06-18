package telegram

import (
	"fmt"
	"strings"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff"
)

func (b *Bot) sendCurrencies(m Message) {
	cur_strings := make([]string, len(tinkoff.Currencies.List))
	for i, cur := range tinkoff.Currencies.List {
		cur_strings[i] = fmt.Sprintf("%v: %v.%v", cur.Name, cur.Units, tinkoff.SCurNano(cur.Nano))
	}
	msg := strings.Join(cur_strings, "\n")
	b.SendMessage(m.Chat.ID, msg)
}
