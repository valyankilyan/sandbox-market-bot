package telegram

import (
	"fmt"
	"log"
	"strconv"
)

func (b *TBot) tinkoffToken(m Message, cmd []string) {
	if len(cmd) == 1 {
		err := b.server.UpdateTinkoffToken(m.From, cmd[0])
		if err != nil {
			log.Println("cant't update token", err)
			b.sendError(m.Chat.ID, "Что-то пошло не так...")
		} else {
			b.sendMessage(m.Chat.ID, "Токен обновлен")
		}
	} else {
		b.sendError(m.Chat.ID, "Слишком мало или слишком много аргументов.")
	}

}

func (b *TBot) payIn(m Message, cmd []string) {
	if len(cmd) == 0 {
		b.sendError(m.Chat.ID, "Кажется, слишком мало аргументов")
	}
	if len(cmd) > 2 {
		b.sendError(m.Chat.ID, "Кажется слишком много аргументов")
	}

	var units int64
	var nano int32
	units, err := strconv.ParseInt(cmd[0], 10, 64)
	if err != nil {
		b.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на целое число.", cmd[1]))
	}
	if len(cmd) == 2 {
		if nano64, err := strconv.ParseInt(cmd[1], 10, 32); err != nil {
			b.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на целое число.", cmd[2]))
		} else {
			nano = int32(nano64)
		}
	}

	tt, err := b.server.UserToken(User{TgId: m.From.TgId})
	if err != nil {
		b.sendError(m.Chat.ID, "Что-то пошло не так...")
		return
	}
	fmt.Println(units, nano, tt)
	balance, err := b.invest.PayIn(tt, Quotation{
		Units: units,
		Nano:  nano,
	})
	if err != nil {
		b.sendError(m.Chat.ID, "Что-то пошло не так...")
		log.Printf("%v while paying in", err)
		return
	}
	b.sendMessage(m.Chat.ID,
		fmt.Sprintf("У вас %v рублей на счету.", formatQuotation(balance)))
}
