package telegram

import (
	"fmt"
	"strings"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff"
)

var currencyFound string = "Вот что мне удалось найти:"

var currencyNotFound string = `Ничего найти не получилось. 
Попробуй /currencies, чтобы узнать короткие имена валют.`

func formatCurrency(cur tinkoff.Currency) string {
	return fmt.Sprintf("%v: %v.%v", cur.Name, cur.Units, tinkoff.SCurNano(cur.Nano))
}

func (b *Bot) sendCurrencies(m Message) {
	cur_strings := make([]string, len(tinkoff.Currencies.List))
	for i, cur := range tinkoff.Currencies.List {
		cur_strings[i] = fmt.Sprintf("(%v) %v", cur.Shortname, formatCurrency(cur))
	}
	msg := strings.Join(cur_strings, "\n")
	b.SendMessage(m.Chat.ID, msg)
}

func (b *Bot) sendCurrency(m Message, currencies []string) {
	msg := []string{currencyFound}
	for _, cur := range currencies {
		c, found := tinkoff.GetCurrency(cur)
		if found {
			msg = append(msg, formatCurrency(c))
		}
	}
	if len(msg) == 1 {
		b.SendMessage(m.Chat.ID, currencyNotFound)
	} else {
		b.SendMessage(m.Chat.ID, strings.Join(msg, "\n"))
	}
}
