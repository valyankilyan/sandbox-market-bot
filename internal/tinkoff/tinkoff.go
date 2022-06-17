package tinkoff

type Tinkoff struct {
	Token string

	Currencies     []Currency
	CurrenciesFigi []string
	// Conn  *grpc.ClientConn
}

func New(token string) *Tinkoff {
	return &Tinkoff{
		Token: token,
	}
}

type Currency struct {
	name      string
	shortname string
	figi      string
	units     int64
	nano      int32
}

type MoneyValue struct {
	Currency string
	units    int64
	nano     int32
}
