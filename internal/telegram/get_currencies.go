package telegram

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff"
)

var currencyFound string = "Вот что мне удалось найти:"

var currencyNotFound string = `Ничего найти не получилось. 
Попробуй /currencies, чтобы узнать короткие имена валют.`

func formatCurrency(cur tinkoff.Currency) string {
	return fmt.Sprintf("%v: %v", strings.ToUpper(cur.Shortname), tinkoff.SCurPrice(cur.Units, cur.Nano))
	// return fmt.Sprintf("%v: %v.%v", cur.Name, cur.Units, tinkoff.SCurNano(cur.Nano))
}

func (b *Bot) sendCurrencies(m Message) {
	cur_strings := make([]string, len(tinkoff.Currencies.List))
	for i, cur := range tinkoff.Currencies.List {
		cur_strings[i] = fmt.Sprintf("%v", formatCurrency(cur))
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

func (b *Bot) buyCurrency(m Message, cmd []string) {
	if len(cmd) <= 1 {
		b.sendError(m.Chat.ID, "Слишком мало аргументов")
		return
	}

	cur, done := tinkoff.GetCurrency(cmd[0])
	if !done {
		b.sendError(m.Chat.ID, fmt.Sprintf("Валюта %v не распознана", cmd[0]))
		return
	}

	quant, err := strconv.ParseInt(cmd[1], 10, 64)
	if err != nil || quant <= 0 {
		b.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на натуральное число.", cmd[1]))
		return
	}

	tt, err := b.readUserToken(m.From.ID)
	if err != nil {
		b.sendError(m.Chat.ID, "Может быть у вас не добавлен токен?")
		return
	}

	done, amount, err := tinkoff.New(tt).BuyCurrencyMarket(cur, quant)
	if err != nil {
		fmt.Printf("Error while buying currency %v\n", err)
	}

	if done {
		b.SendMessage(m.Chat.ID, fmt.Sprintf("Транзакция прошла успешно. Потрачено %v.%v рублей.", amount.Units, tinkoff.SCurNano(amount.Nano)))
	} else {
		b.SendMessage(m.Chat.ID, "Ошбика транзакции.")
	}
}
