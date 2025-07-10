package user

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserDetailRoutes(router fiber.Router, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userDetailRepo := repositories.NewUserDetailRepository(db)
	userDetailService := services.NewUserDetailService(userDetailRepo, userRepo)
	userDetailHandler := handlers.NewUserDetailHandler(userDetailService)

	userDetail := router.Group("/user-details")
	userDetail.Use(middlewares.JWTMiddleware)

	userDetail.Post("/create", middlewares.RequirePermission("user_detail.create"), userDetailHandler.Create)
	userDetail.Get("/get/:uuid", middlewares.RequirePermission("user_detail.view"), userDetailHandler.Get)
	userDetail.Patch("/update/:uuid", middlewares.RequirePermission("user_detail.update"), userDetailHandler.Update)
	userDetail.Delete("/delete/:uuid", middlewares.RequirePermission("user_detail.delete"), userDetailHandler.Delete)
}
