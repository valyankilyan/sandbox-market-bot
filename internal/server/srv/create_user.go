package srv

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
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
