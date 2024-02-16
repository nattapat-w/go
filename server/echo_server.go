package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nattapat-w/go/config"
	"github.com/nattapat-w/go/order"
	"gorm.io/gorm"
)

type fiberServer struct {
	app *fiber.App
	db  *gorm.DB
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) Server {
	return &fiberServer{
		app: fiber.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *fiberServer) Start() {
	// Initialize module(feature) here
	// ...
	order.InitializeOrderModule(s.app, s.db)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Listen(serverUrl)
}
