package order

import (
	"github.com/gofiber/fiber/v2"
	orderController "github.com/nattapat-w/go/order/controller"
	orderRepositoryDBImpl "github.com/nattapat-w/go/order/repository/impl"
	orderUsecaseImpl "github.com/nattapat-w/go/order/usecases/impl"
	"gorm.io/gorm"
)

func InitializeOrderModule(app *fiber.App, db *gorm.DB) {
	orderRepositoryDB := orderRepositoryDBImpl.NewOrderRepsoitryDBImpl(db)
	orderUsecase := orderUsecaseImpl.NewOrderUsecaesImpl(&orderRepositoryDB)
	orderController := orderController.NewOrderController(&orderUsecase)

	orderRouters := app.Group("/api/v1/order")
	orderRouters.Post("/", orderController.CreateOrder)
}
