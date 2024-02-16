package controller

import (
	"github.com/gofiber/fiber/v2"
	response_model "github.com/nattapat-w/go/model"
	"github.com/nattapat-w/go/order/model"
	"github.com/nattapat-w/go/order/usecases"
)

type OrderController struct {
	orderUsecase usecases.OrderUsecase
}

func NewOrderController(orderUsecase *usecases.OrderUsecase) OrderController {
	return OrderController{orderUsecase: *orderUsecase}
}
func (controller OrderController) Route(app *fiber.App) {
	router := app.Group("/api/v1/order")
	router.Post("/", controller.CreateOrder)
}

func (controller OrderController) CreateOrder(ctx *fiber.Ctx) error {
	var body model.OrderCreateOrUpdateModel
	err := ctx.BodyParser(&body)
	if err != nil {
		panic("invalid request body")
	}

	response := controller.orderUsecase.CreateOrder(ctx.Context(), body)

	return ctx.Status(fiber.StatusCreated).JSON(response_model.GeneralResponseModel{
		Code:    "201",
		Message: "Order Created Success!",
		Data:    response,
	})
}
