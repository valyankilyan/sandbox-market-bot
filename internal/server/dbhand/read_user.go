package dbhand

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
)

func (d *dbhandler) ReadUser(ctx context.Context, Id int) (models.User, error) {
	return models.User{}, nil
}
