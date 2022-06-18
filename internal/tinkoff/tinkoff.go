package tinkoff

import "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"

var defTink *Tinkoff

func DefTinkInit() {
	defTink = New(config.Conf.Tinkoff.DefaultToken)
}

type Tinkoff struct {
	Token string
}

func New(token string) *Tinkoff {
	return &Tinkoff{
		Token: token,
	}
}

type MoneyValue struct {
	Currency string
	Units    int64
	Nano     int32
}

type Quotation struct {
	Units int64
	Nano  int32
}
