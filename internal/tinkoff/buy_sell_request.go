package tinkoff

func (t *Tinkoff) BuyCurrencyMarket(cur Currency, quant int64) (bool, MoneyValue, error) {
	lastPrice := GetLastPrice(cur.Figi)
	return t.PostOrderRequest(PostOrderRequest{
		Figi:      cur.Figi,
		Quantity:  quant,
		Price:     &lastPrice,
		Direction: OrderDirection_ORDER_DIRECTION_BUY,
		OrderType: OrderType_ORDER_TYPE_MARKET,
	})
}

func (t *Tinkoff) SellCurrencyMarket(cur Currency, quant int64) (bool, MoneyValue, error) {
	lastPrice := GetLastPrice(cur.Figi)
	return t.PostOrderRequest(PostOrderRequest{
		Figi:      cur.Figi,
		Quantity:  quant,
		Price:     &lastPrice,
		Direction: OrderDirection_ORDER_DIRECTION_SELL,
		OrderType: OrderType_ORDER_TYPE_MARKET,
	})
}
