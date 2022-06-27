package telegram

import "log"

func (b *TBot) tinkoffToken(m Message, cmd []string) {
	if len(cmd) == 2 {
		err := b.server.UpdateTinkoffToken(m.From, cmd[1])
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

// func (b *Bot) payIn(m Message, cmd []string) {
// 	if len(cmd) == 1 {
// 		b.sendError(m.Chat.ID, "Кажется, слишком мало аргументов")
// 	}
// 	if len(cmd) > 3 {
// 		b.sendError(m.Chat.ID, "Кажется слишком много аргументов")
// 	}

// 	var units int64
// 	var nano int32
// 	units, err := strconv.ParseInt(cmd[1], 10, 64)
// 	if err != nil {
// 		b.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на целое число.", cmd[1]))
// 	}
// 	if len(cmd) == 3 {
// 		if nano64, err := strconv.ParseInt(cmd[2], 10, 32); err != nil {
// 			b.sendError(m.Chat.ID, fmt.Sprintf("%v не похоже на целое число.", cmd[2]))
// 		} else {
// 			nano = int32(nano64)
// 		}
// 	}

// 	tt, err := b.readUserToken(m.From.ID)
// 	if err != nil {
// 		b.sendError(m.Chat.ID, "Что-то пошло не так...")
// 		return
// 	}
// 	units_out, nano_out, err := tinkoff.New(tt).PayIn(units, nano)
// 	if err != nil {
// 		b.sendError(m.Chat.ID, "Что-то пошло не так...")
// 		fmt.Printf("%v while paying in", err)
// 		return
// 	}
// 	b.SendMessage(m.Chat.ID, fmt.Sprintf("У вас %v.%v рублей на счету.", units_out, nano_out))
// }
