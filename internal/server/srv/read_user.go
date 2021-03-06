package srv

import (
	"context"
	"errors"
	"fmt"

	"github.com/valyankilyan/sandbox-market-bot/internal/server/models"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (t *tserver) ReadUser(ctx context.Context, tgid *pb.TgId) (*pb.User, error) {
	var usr models.User
	err := t.db.Where("tg_id = ?", fmt.Sprint(tgid.TgId)).First(&usr).Error
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
