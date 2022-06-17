package srv

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
)

type Repository interface {
	CreateUser(context.Context, models.User) (int, error)
	ReadUser(context.Context, int) (models.User, error)
	UpdateUser(context.Context, models.User) error
	DeleteUser(context.Context, int) error
}
