package telegram

const start = `
С помощью этого бота можно играть в песочнице тинькофф инвестиций.
Внесите свой токен и посмотрите, что вы сможете с этим сделать /help`

const help = `
/start - старт и регистрация
/help - помощь
/tinkoff_token <token> - добавление токена песочницы`

func (b *Bot) sendStart(chat_id int64) {
	b.SendMessage(chat_id, start)
}

func (b *Bot) sendHelp(chat_id int64) {
	b.SendMessage(chat_id, help)
}

func (b *Bot) notRecognized(chat_id int64) {
	b.SendMessage(chat_id, "Извини, я не понял, используй /help")
}
