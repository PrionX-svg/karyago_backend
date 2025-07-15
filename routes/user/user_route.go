package user

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(router fiber.Router, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepositories(db)
	branchRepo := repositories.NewBranchRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	userService := services.NewUserService(db, userRepo, roleRepo, branchRepo, employeeRepo, companyRepo)
	userHandler := handlers.NewUserHandler(userService)

	user := router.Group("/users")
	user.Use(middlewares.JWTMiddleware)

	user.Post("/create", middlewares.RequirePermission("user.create"), userHandler.CreateUser)
	user.Get("/@me", userHandler.GetMe)
	user.Get("/get-with-emp", middlewares.RequirePermission("user.view-all"), userHandler.GetAllUsers)
	user.Get("/get-all/dt", middlewares.RequirePermission("user.view-all-datatable"), userHandler.GetUsersDataTable)
	user.Get("/get/:uuid", middlewares.RequirePermission("user.view"), userHandler.GetUserByUUID)
	user.Delete("/delete/:uuid", middlewares.RequirePermission("user.delete"), userHandler.DeleteUser)
	user.Patch("/update/:uuid", middlewares.RequirePermission("user.update"), userHandler.UpdateUser)
	user.Patch("/rehire/:uuid", middlewares.RequirePermission("user.rehire"), userHandler.RehireUser)
}
