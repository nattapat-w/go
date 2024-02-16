package impl

import (
	"context"

	"github.com/nattapat-w/go/order/entities"
	"github.com/nattapat-w/go/order/model"
	"github.com/nattapat-w/go/order/repository"
	service "github.com/nattapat-w/go/order/usecases"
)

type orderUsecaseImpl struct {
	orderRepository repository.OrderRepository
}

func NewOrderUsecaesImpl(orderRepository *repository.OrderRepository) service.OrderUsecase {
	return orderUsecaseImpl{orderRepository: *orderRepository}
}
func (service orderUsecaseImpl) CreateOrder(ctx context.Context, request model.OrderCreateOrUpdateModel) *model.OrderCreateOrUpdateModel {
	order := entities.Order{
		Total: request.Total,
	}
	orderId := service.orderRepository.Create(ctx, order)
	return &model.OrderCreateOrUpdateModel{
		ID:    orderId,
		Total: order.Total,
	}
}

// func (service orderServieImpl) FindOrder(ctx context.Context, id uint) *model.OrderModel {
// 	return &model.OrderModel{
// 		ID:    orderId,
// 		Total: order.Total,
// 	}
// }
