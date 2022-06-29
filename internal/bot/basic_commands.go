package bot

import "fmt"

const start = `
С помощью этого бота можно играть в песочнице тинькофф инвестиций.
Внесите свой токен и посмотрите, что вы сможете с этим сделать /help`

const basic_help = `
/start - старт и регистрация
/help - помощь
/tinkoff_token <token> - добавление токена песочницы
/payin <units (uint)> <nano (uint optional)> - добавление рублей себе на счет (может быть отрицательным числом).
Units - это рубли, nano - это копейки`

// /balance - ваш баланс в валютах`

const currencyHelp = `
/currency <shortname>... - стоимость валют`

var help string

func init() {
	help = fmt.Sprintf("%v\n%v", basic_help, currencyHelp)
}

func (bs *BotService) sendStart(chat_id int64) {
	bs.bot.SendMessage(chat_id, start)
}

func (bs *BotService) sendHelp(chat_id int64) {
	bs.bot.SendMessage(chat_id, help)
}

func (bs *BotService) notRecognized(chat_id int64) {
	bs.bot.SendMessage(chat_id, "Извини, я не понял, используй /help")
}

func (bs *BotService) sendError(chat_id int64, err string) {
	bs.bot.SendMessage(chat_id, err)
}
