package srv

import (
	"context"

	"github.com/valyankilyan/sandbox-market-bot/internal/server/models"
)

type Repository interface {
	CreateUser(context.Context, models.User) (int, error)
	ReadUser(context.Context, int) (models.User, error)
	UpdateUser(context.Context, models.User) error
	DeleteUser(context.Context, int) error
}
