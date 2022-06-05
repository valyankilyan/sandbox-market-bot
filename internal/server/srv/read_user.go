package srv

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
)

func (t *tserver) ReadUser(ctx context.Context, usr *api.Id) (*api.User, error) {
	return &api.User{}, nil
}
