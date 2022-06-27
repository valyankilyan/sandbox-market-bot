package telegram

import (
	"fmt"
	"strings"
)

func (b *TBot) currencies(chat_id int64, cmd []string) {
	var curs []Currency
	var err error

	if len(cmd) == 0 {
		curs, err = b.invest.AllCurrencies()
	} else {
		curs, err = b.invest.Currencies(cmd)
	}

	if err != nil {
		b.sendError(chat_id, "Что-то не получается...")
		return
	}
	b.sendMessage(chat_id, formatCurrencies(curs))
}

func formatCurrencies(curs []Currency) string {
	ans := make([]string, len(curs))
	for i, c := range curs {
		ans[i] = fmt.Sprintf("%v - (%v): %v Lot:%v",
			formatQuotation(c.Price), strings.ToUpper(c.ShortName), c.Name, c.Lot)
	}
	return strings.Join(ans, "\n")
}

func formatQuotation(q Quotation) string {
	if q.Nano != 0 {
		return fmt.Sprintf("%v.%v", q.Units, formatCurNano(q.Nano))
	} else {
		return fmt.Sprintf("%v", q.Units)
	}
}

func formatCurNano(nano int32) string {
	ans := fmt.Sprintf("%d", nano)
	ans = strings.TrimRight(ans, "0")
	return ans
}
