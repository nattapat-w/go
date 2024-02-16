package usecases

import (
	"context"

	"github.com/nattapat-w/go/order/model"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, request model.OrderCreateOrUpdateModel) (response *model.OrderCreateOrUpdateModel)
	// FindOrder(ctx context.Context, id uint) (response *model.OrderModel)
}
