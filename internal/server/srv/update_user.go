package srv

import (
	"context"
	"errors"

	"github.com/valyankilyan/sandbox-market-bot/internal/server/models"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func (t *tserver) UpdateUser(ctx context.Context, usr *pb.User) (*emptypb.Empty, error) {
	upusr := models.User{
		Model:        gorm.Model{ID: uint(usr.Id)},
		Id:           usr.Id,
		Name:         usr.Name,
		TgUserName:   usr.TgUserName,
		TgId:         usr.TgId,
		TinkoffToken: usr.TinkoffToken,
	}

	err := t.db.Model(&upusr).Updates(&upusr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &emptypb.Empty{}, nil
}
