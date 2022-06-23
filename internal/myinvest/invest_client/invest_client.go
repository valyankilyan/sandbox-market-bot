package invest_client

import "github.com/valyankilyan/sandbox-market-bot/config"

type InvestClient struct {
	Token string
}

func DefaultInvestClientInit() *InvestClient {
	return New(config.Tinkoff.DefaultToken)
}

func New(token string) *InvestClient {
	return &InvestClient{Token: token}
}
