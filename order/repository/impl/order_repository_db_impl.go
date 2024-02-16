package impl

import (
	"context"

	"github.com/nattapat-w/go/order/entities"
	"github.com/nattapat-w/go/order/repository"
	"gorm.io/gorm"
)

type orderRepositoryDBImpl struct {
	db *gorm.DB
}

func NewOrderRepsoitryDBImpl(db *gorm.DB) repository.OrderRepository {
	err := db.AutoMigrate(&entities.Order{})
	if err != nil {
		panic(err)
	}
	return orderRepositoryDBImpl{db: db}
}
func (repository orderRepositoryDBImpl) Create(ctx context.Context, order entities.Order) (id uint) {
	err := repository.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		panic(err)
	}
	return order.ID
}
