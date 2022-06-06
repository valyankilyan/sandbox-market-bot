package srv

import (
	"context"
	"errors"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (t *tserver) ReadUser(ctx context.Context, usrid *pb.Id) (*pb.User, error) {
	var usr models.User
	err := t.db.First(&usr, usrid.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &pb.User{
		Id:           usr.Id,
		Name:         usr.Name,
		TgUserName:   usr.TgUserName,
		TgId:         usr.TgId,
		TinkoffToken: usr.TinkoffToken,
	}, err
}
