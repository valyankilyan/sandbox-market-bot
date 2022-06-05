package dbhand

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")

type dbhandler struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *dbhandler {
	return &dbhandler{pool: pool}
}
