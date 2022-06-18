package tinkoff

type Tinkoff struct {
	Token string

	// Conn  *grpc.ClientConn
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
