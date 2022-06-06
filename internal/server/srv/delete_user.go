package srv

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) DeleteUser(ctx context.Context, Id *pb.Id) (*emptypb.Empty, error) {

	err := t.db.Delete(&models.User{}, Id.Id).Error

	return &emptypb.Empty{}, err
}
