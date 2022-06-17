package telegram

import (
	"fmt"
	"strconv"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff"
)

func (b *Bot) payIn(m Message, cmd []string) {
	if len(cmd) == 1 {
		b.sendError(m.Chat.ID, "Кажется, слишком мало аргументов")
	}
	if len(cmd) > 3 {
		b.sendError(m.Chat.ID, "Кажется слишком много аргументов")
	}

	var units, nano uint64
	units, err := strconv.ParseUint(cmd[1], 10, 64)
	if err != nil {
		b.sendError(m.Chat.ID, "Units не похоже на натуральное число")
	}
	if len(cmd) == 3 {
		if nano, err = strconv.ParseUint(cmd[2], 10, 64); err != nil {
			b.sendError(m.Chat.ID, "Nano не похоже на натуральное число")
		}
	}

	tt, err := b.readUserToken(m.From.ID)
	if err != nil {
		b.sendError(m.Chat.ID, "Что-то пошло не так...")
		return
	}
	units_out, nano_out, err := tinkoff.New(tt).PayIn(int64(units), int32(nano))
	if err != nil {
		b.sendError(m.Chat.ID, "Что-то пошло не так...")
		return
	}
	b.SendMessage(m.Chat.ID, fmt.Sprintf("У вас %v.%v рублей на счету.", units_out, nano_out))
}

func (b *Bot) tinkoffToken(m Message, cmd []string) {
	if len(cmd) == 2 {
		err := b.updateTinkoffToken(m.From.ID, cmd[1])
		if err != nil {
			b.sendError(m.Chat.ID, "Что-то пошло не так...")
		} else {
			b.SendMessage(m.Chat.ID, "Токен обновлен")
		}
	} else {
		b.sendError(m.Chat.ID, "Слишком мало или слишком много аргументов.")
	}
}
