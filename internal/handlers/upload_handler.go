package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/services"
)

type UploadHandler struct {
	UploadService services.UploadService
}

func NewUploadHandler(svc services.UploadService) *UploadHandler {
	return &UploadHandler{UploadService: svc}
}

func (h *UploadHandler) UploadImage(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	url, err := h.UploadService.UploadImage(fileHeader, "uploads")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"url": url,
	})
}

func (h *UploadHandler) DeleteImage(c *fiber.Ctx) error {
	var payload struct {
		FileName string `json:"file_name"`
	}

	if err := c.BodyParser(&payload); err != nil || payload.FileName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file_name is required",
		})
	}

	if err := h.UploadService.DeleteImage(payload.FileName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
