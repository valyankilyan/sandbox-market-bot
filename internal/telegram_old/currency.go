package telegram_old

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff"
)

var currencyFound string = "Вот что мне удалось найти:"

var currencyNotFound string = `Ничего найти не получилось. 
Попробуй /currency list, чтобы узнать короткие имена валют.`

var currencyHelp string = `	/currency 
	list - посмотреть курс всех доступных валют
	list <name>... - посмотреть курс конкректных валют
	buy <name> n - купить n лотов валюты
	sell <name> n - продать n лотов валюты`

func formatCurrency(cur tinkoff.Currency) string {
	return fmt.Sprintf("%v: %v", strings.ToUpper(cur.Shortname),
		tinkoff.SCurPrice(tinkoff.Quotation{Units: cur.Units, Nano: cur.Nano}))
}

func (b *Bot) currency(m Message, cmd []string) {
	if len(cmd) == 0 {
		b.SendMessage(m.Chat.ID, currencyHelp)
		return
	}

	switch cmd[0] {
	case "list":
		b.currencyList(m, cmd[1:])
	case "buy":
		b.currencyBuy(m, cmd[1:])
	case "sell":
		b.currencySell(m, cmd[1:])
	default:
		b.SendMessage(m.Chat.ID, currencyHelp)
	}

}

func (b *Bot) currencyList(m Message, cmd []string) {
	if len(cmd) == 0 {
		b.sendAllCurrencies(m)
	} else {
		b.sendCurrencies(m, cmd)
	}
}

func (b *Bot) sendAllCurrencies(m Message) {
	cur_strings := make([]string, len(tinkoff.Currencies.List))
	for i, cur := range tinkoff.Currencies.List {
		cur_strings[i] = fmt.Sprintf("%v, L%v", formatCurrency(cur), cur.Lot)
	}
	msg := strings.Join(cur_strings, "\n")
	b.SendMessage(m.Chat.ID, msg)
}

func (b *Bot) sendCurrencies(m Message, currencies []string) {
	msg := []string{currencyFound}
	for _, cur := range currencies {
		c, found := tinkoff.GetCurrency(cur)
		if found {
			msg = append(msg, fmt.Sprintf("%v, L%v", formatCurrency(c), c.Lot))
		}
	}
	if len(msg) == 1 {
		b.SendMessage(m.Chat.ID, currencyNotFound)
	} else {
		b.SendMessage(m.Chat.ID, strings.Join(msg, "\n"))
	}
}

func (b *Bot) currencyBuy(m Message, cmd []string) {
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
		b.SendMessage(m.Chat.ID, fmt.Sprintf("Транзакция прошла успешно. Потрачено %v рублей.",
			tinkoff.SCurPrice(tinkoff.Quotation{Units: amount.Units, Nano: amount.Nano})))
	} else {
		b.SendMessage(m.Chat.ID, "Ошбика транзакции.")
	}
}

func (b *Bot) currencySell(m Message, cmd []string) {
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

	done, _, err = tinkoff.New(tt).SellCurrencyMarket(cur, quant)
	if err != nil {
		fmt.Printf("Error while buying currency %v\n", err)
	}

	if done {
		b.SendMessage(m.Chat.ID, "Транзакция прошла успешно.")
	} else {
		b.SendMessage(m.Chat.ID, "Ошбика транзакции.")
	}
}