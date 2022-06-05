package srv

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) DeleteUser(ctx context.Context, Id *api.Id) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
