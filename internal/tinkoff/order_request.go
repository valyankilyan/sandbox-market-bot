package tinkoff

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff/investapi"
)

type OrderDirection int32

const (
	OrderDirection_ORDER_DIRECTION_UNSPECIFIED OrderDirection = 0 //Значение не указано
	OrderDirection_ORDER_DIRECTION_BUY         OrderDirection = 1 //Покупка
	OrderDirection_ORDER_DIRECTION_SELL        OrderDirection = 2 //Продажа
)

type OrderType int32

const (
	OrderType_ORDER_TYPE_UNSPECIFIED OrderType = 0 //Значение не указано
	OrderType_ORDER_TYPE_LIMIT       OrderType = 1 //Лимитная
	OrderType_ORDER_TYPE_MARKET      OrderType = 2 //Рыночная
)

type PostOrderRequest struct {
	Figi      string
	Quantity  int64
	Price     *Quotation
	Direction OrderDirection
	// AccountId string
	OrderType OrderType
	// OrderId   string
}

func (t *Tinkoff) PostOrderRequest(poreq PostOrderRequest) (bool, MoneyValue, error) {
	ctx, cancel := t.getContext()
	defer cancel()

	conn, err := t.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return false, MoneyValue{}, err
	}
	defer conn.Close()

	client := investapi.NewSandboxServiceClient(conn)
	account_id, err := t.sandboxAccountId(ctx, client)

	invest_poreq := investapi.PostOrderRequest{
		Figi:     poreq.Figi,
		Quantity: poreq.Quantity,
		Price: &investapi.Quotation{
			Units: poreq.Price.Units,
			Nano:  poreq.Price.Nano,
		},
		Direction: investapi.OrderDirection(poreq.Direction),
		OrderType: investapi.OrderType(poreq.OrderType),
		AccountId: account_id,
		OrderId:   randStringBytes(36),
	}

	invest_poresp, err := client.PostSandboxOrder(ctx, &invest_poreq)
	if err != nil {
		fmt.Println(err.Error())
		return false, MoneyValue{}, err
	}

	fmt.Println(invest_poresp)

	amount := MoneyValue{
		Currency: invest_poresp.TotalOrderAmount.Currency,
		Units:    invest_poresp.TotalOrderAmount.Units,
		Nano:     invest_poresp.TotalOrderAmount.Nano,
	}
	return true, amount, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
