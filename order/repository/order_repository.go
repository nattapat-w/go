package repository

import (
	"context"

	"github.com/nattapat-w/go/order/entities"
)

type OrderRepository interface {
	Create(ctx context.Context, order entities.Order) (id uint)
}
