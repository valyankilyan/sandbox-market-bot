package bot

import (
	"fmt"
	"log"
	"strconv"
)

func (bs *BotService) tinkoffToken(m Message, cmd []string) {
	if len(cmd) == 1 {
		err := bs.server.UpdateTinkoffToken(m.From, cmd[0])
		if err != nil {
			log.Println("cant't update token", err)
			bs.sendError(m.Chat.ID, "Что-то пошло не так...")
		} else {
			bs.bot.SendMessage(m.Chat.ID, "Токен обновлен")
		}
	} else {
		bs.sendError(m.Chat.ID, "Слишком мало или слишком много аргументов.")
	}

}

func (bs *BotService) payIn(m Message, cmd []string) {
	if len(cmd) == 0 {
		bs.sendError(m.Chat.ID, "Кажется, слишком мало аргументов")
	}
	if len(cmd) > 2 {
		bs.sendError(m.Chat.ID, "Кажется слишком много аргументов")
	}

	var units int64
	var nano int32
	units, err := strconv.ParseInt(cmd[0], 10, 64)
	if err != nil {
		bs.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на целое число.", cmd[1]))
	}
	if len(cmd) == 2 {
		if nano64, err := strconv.ParseInt(cmd[1], 10, 32); err != nil {
			bs.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на целое число.", cmd[2]))
		} else {
			nano = int32(nano64)
		}
	}

	tt, err := bs.server.UserToken(User{TgId: m.From.TgId})
	if err != nil {
		bs.sendError(m.Chat.ID, "Что-то пошло не так...")
		return
	}
	fmt.Println(units, nano, tt)
	balance, err := bs.invest.PayIn(tt, Quotation{
		Units: units,
		Nano:  nano,
	})
	if err != nil {
		bs.sendError(m.Chat.ID, "Что-то пошло не так...")
		log.Printf("%v while paying in", err)
		return
	}
	bs.bot.SendMessage(m.Chat.ID,
		fmt.Sprintf("У вас %v рублей на счету.", formatQuotation(balance)))
}
