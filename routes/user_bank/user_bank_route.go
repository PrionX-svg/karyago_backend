package usereducation

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserBankRoutes (router fiber.Router, db *gorm.DB) {
	userBankRepository := repositories.NewUserBankRepository(db)
	userRepository := repositories.NewUserRepository(db)
	userBankService := services.NewUserBankService(userBankRepository, userRepository)
	UserBankHandler := handlers.NewUserBankHandler(userBankService)

	userBanks := router.Group("/user-banks")
	userBanks.Use(middlewares.JWTMiddleware)

	userBanks.Post("/create", middlewares.RequirePermission("user-banks.create"), UserBankHandler.Create)
	userBanks.Get("/get-all", middlewares.RequirePermission("user-banks.view-all"), UserBankHandler.GetAll)
	userBanks.Get("/get-by-id/:id", middlewares.RequirePermission("user-banks.view"), UserBankHandler.GetByID)
	userBanks.Get("/get-by-uuid/:uuid", middlewares.RequirePermission("user-banks.view"), UserBankHandler.GetByUUID)
	userBanks.Get("/get-by-user/:uuid", middlewares.RequirePermission("user-banks.view"), UserBankHandler.GetByUserUUID)
	userBanks.Patch("/update/:uuid", middlewares.RequirePermission("user-banks.update"), UserBankHandler.Update)
	userBanks.Delete("/delete/:uuid", middlewares.RequirePermission("user-banks.delete"), UserBankHandler.Delete)
}