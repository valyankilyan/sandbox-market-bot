package bot

import "strings"

type BotService struct {
	bot    Bot
	server Server
	invest Invest
}

func New(bot Bot, server Server, invest Invest) *BotService {
	return &BotService{
		bot:    bot,
		server: server,
		invest: invest,
	}
}

func (bs *BotService) MessageProcessor() {
	msgch := make(chan Message, 20)
	go bs.bot.GetUpdates(msgch)
	for {
		msg := <-msgch

		text := strings.Split(msg.Text, " ")
		if len(text) == 0 {
			go bs.notRecognized(msg.Chat.ID)
		}

		switch text[0] {
		case "/start":
			go bs.sendStart(msg.Chat.ID)
			go bs.server.CreateUser(msg.From)
		case "/help":
			go bs.sendHelp(msg.Chat.ID)
		case "/tinkoff_token":
			go bs.tinkoffToken(msg, text[1:])
		case "/payin":
			go bs.payIn(msg, text[1:])
		case "/currency":
			go bs.currencies(msg.Chat.ID, text[1:])
		// case "/balance":
		// 	bs.balance(m)
		default:
			go bs.notRecognized(msg.Chat.ID)
		}
	}

}

type Bot interface {
	GetUpdates(msgch chan Message)
	SendMessage(chatId int64, message string) error
}
type Server interface {
	CreateUser(user User)
	UpdateTinkoffToken(user User, token string) error
	UserToken(user User) (token string, err error)
}

type Invest interface {
	AllCurrencies() ([]Currency, error)
	Currencies(shortnames []string) ([]Currency, error)
	PayIn(token string, quantity Quotation) (balance Quotation, err error)
}
