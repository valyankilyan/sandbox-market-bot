package srv

import (
	"context"

	"github.com/valyankilyan/sandbox-market-bot/internal/server/models"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/api"
)

func (t *tserver) CreateUser(ctx context.Context, usr *pb.User) (*pb.Id, error) {
	newuser := models.User{
		Name:         usr.Name,
		TgUserName:   usr.TgUserName,
		TgId:         usr.TgId,
		TinkoffToken: usr.TinkoffToken,
	}

	t.db.Create(&newuser)

	return &pb.Id{
		Id: newuser.Id,
	}, nil
}
