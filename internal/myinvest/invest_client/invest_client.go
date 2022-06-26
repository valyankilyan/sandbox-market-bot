package invest_client

import "log"

type InvestClient struct{}

func NewInvestClient() *InvestClient {
	log.Println("newinvestclient")
	return &InvestClient{}
}
