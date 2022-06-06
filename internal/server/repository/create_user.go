package repository

import (
	"context"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
)

func (r *repository) CreateUser(ctx context.Context, usr models.User) (Id int, err error) {
	// span, ctx := opentracing.StartSpanFromContext(ctx, "db")

	// defer span.Finish()

	// const query = `
	// 	insert into  (
	// 		name,
	// 		"desc",
	// 		price,
	// 		created_at
	// 	) VALUES (
	// 		$1, $2, $3, now()
	// 	) returning id
	// `

	// err = r.pool.QueryRow(ctx, query,
	// 	product.Name,
	// 	product.Desc,
	// 	product.Price,
	// ).Scan(&ID)

	return 0, nil
}
