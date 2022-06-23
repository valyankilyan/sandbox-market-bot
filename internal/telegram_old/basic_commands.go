package telegram_old

import "fmt"

const start = `
С помощью этого бота можно играть в песочнице тинькофф инвестиций.
Внесите свой токен и посмотрите, что вы сможете с этим сделать /help`

const basic_help = `
/start - старт и регистрация
/help - помощь
/tinkoff_token <token> - добавление токена песочницы
/payin <units (uint)> <nano (uint optional)> - добавление рублей себе на счет (может быть отрицательным числом).
Units - это рубли, nano - это копейки
/balance - ваш баланс в валютах`

var help string

func init() {
	help = fmt.Sprintf("%v\n%v", basic_help, currencyHelp)
}

func (b *Bot) sendStart(chat_id int64) {
	b.SendMessage(chat_id, start)
}

func (b *Bot) sendHelp(chat_id int64) {
	b.SendMessage(chat_id, help)
}

func (b *Bot) notRecognized(chat_id int64) {
	b.SendMessage(chat_id, "Извини, я не понял, используй /help")
}

func (b *Bot) sendError(chat_id int64, err string) {
	b.SendMessage(chat_id, err)
}