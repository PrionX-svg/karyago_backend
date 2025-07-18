package upload

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/services"
	"hris_backend/pkg/r2"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

func SetupUploadRoutes(router fiber.Router, r2Client *minio.Client, bucketName string) {
	uploader := r2.NewUploader(r2Client, bucketName)
	uploadService := services.NewUploadService(uploader)
	uploadHandler := handlers.NewUploadHandler(uploadService)

	upload := router.Group("/upload")
	upload.Use(middlewares.JWTMiddleware)

	upload.Post("/", middlewares.RequirePermission("upload.create"), uploadHandler.UploadImage)
	upload.Post("/delete", middlewares.RequirePermission("upload.delete"), uploadHandler.DeleteImage)
}
