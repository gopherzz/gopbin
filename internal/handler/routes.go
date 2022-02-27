package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gopherzz/gopbin/internal/models"
)

func (h *Handler) createDoc(c *fiber.Ctx) error {
	doc := &models.DocumentRequest{}
	if err := c.BodyParser(doc); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	docId, err := h.services.CreateDocument(doc)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": docId})
}

func (h *Handler) getDoc(c *fiber.Ctx) error {
	id := c.Params("id", "000000")

	doc, err := h.services.ReadDocument(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusFound).JSON(doc)
}
