package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gopherzz/gopbin/internal/services"
)

type Handler struct {
	services *services.Services
}

func (h *Handler) Listen(port string) error {
	router := fiber.New()
	router.Use(logger.New())

	api := router.Group("/api")
	v1 := api.Group("/v1")
	doc := v1.Group("/doc")

	doc.Post("/create", h.createDoc)
	doc.Get("/get/:id", h.getDoc)

	return router.Listen(port)
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}
