package telegram

const help = `
/help - помощь`

func (b *Bot) sendHelp(chat_id int64) {
	b.SendMessage(chat_id, help)
}

func (b *Bot) notRecognized(chat_id int64) {
	b.SendMessage(chat_id, "Извини, я не понял, используй /help")
}
