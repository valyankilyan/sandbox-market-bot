package srv

import (
	"context"

	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
)

func (t *tserver) ReadUser(ctx context.Context, usr *pb.Id) (*pb.User, error) {
	return &pb.User{}, nil
}
