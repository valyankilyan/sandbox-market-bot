package srv

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
)

func (t *tserver) CreateUser(ctx context.Context, usr *api.User) (*api.Id, error) {
	return &api.Id{}, nil
}
