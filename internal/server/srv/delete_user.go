package srv

import (
	"context"
	"errors"
	"fmt"

	"github.com/valyankilyan/sandbox-market-bot/internal/server/models"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func (t *tserver) DeleteUser(ctx context.Context, tgid *pb.TgId) (*emptypb.Empty, error) {
	var usr models.User
	err := t.db.Where("tg_id = ?", fmt.Sprint(tgid.TgId)).First(&usr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	err = t.db.Delete(&models.User{}, usr.Id).Error

	return &emptypb.Empty{}, err
}
