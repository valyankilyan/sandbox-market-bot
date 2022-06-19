package telegram

import (
	"fmt"
	"strings"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff"
)

func (b *Bot) balance(m Message) {
	tt, err := b.readUserToken(m.From.ID)
	if err != nil {
		b.sendError(m.Chat.ID, "Кажется, у вас не указан токен...")
		return
	}
	tink := tinkoff.New(tt)
	var balance tinkoff.Balance
	tink.GetSandboxPortfolio(&balance)
	tink.GetSandboxPosition(&balance)

	msg := make([]string, 1)
	msg[0] = fmt.Sprintf("На вашем счету %v rub", tinkoff.SCurPrice(balance.TotalAmount))
	for _, cur := range balance.Currencies {
		msg = append(msg, fmt.Sprintf("%v: %v", strings.ToUpper(cur.Currency),
			tinkoff.SCurPrice(tinkoff.Quotation{Units: cur.Units, Nano: cur.Nano})))
	}
	b.SendMessage(m.Chat.ID, strings.Join(msg, "\n"))
}