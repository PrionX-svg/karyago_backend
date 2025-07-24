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
	otpRepo := repositories.NewOTPRepositories(db)
	roleRepo := repositories.NewRoleRepositories(db)
	branchRepo := repositories.NewBranchRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)
	departmentRepo := repositories.NewDepartmentRepositories(db)

	userService := services.NewUserService(db, userRepo, otpRepo, roleRepo, branchRepo, employeeRepo, companyRepo, departmentRepo)
	employmentHistoryService := services.NewEmploymentHistoryService(db, repositories.NewEmploymentHistoryRepository(db), employeeRepo, roleRepo, companyRepo, branchRepo)
	userExcelService := services.NewUserExcelService(userService, companyRepo, roleRepo, branchRepo, userRepo, employeeRepo, employmentHistoryService)
	userHandler := handlers.NewUserHandler(userService, userExcelService)

	user := router.Group("/users")
	user.Use(middlewares.JWTMiddleware)

	user.Post("/create", middlewares.RequirePermission("user.create"), userHandler.CreateUser)
	user.Get("/@me", userHandler.GetMe)
	user.Get("/get-with-emp", middlewares.RequirePermission("user.view-all"), userHandler.GetAllUsers)
	user.Get("/get-all/dt", middlewares.RequirePermission("user.view-all-datatable"), userHandler.GetUsersDataTable)
	user.Get("/get/:uuid", middlewares.RequirePermission("user.view"), userHandler.GetUserByUUID)
	user.Delete("/delete/:uuid", middlewares.RequirePermission("user.delete"), userHandler.DeleteUser)
	user.Patch("/update/:uuid", middlewares.RequirePermission("user.update"), userHandler.UpdateUser)
	user.Patch("/update-department/:uuid", middlewares.RequirePermission("user.update-department"), userHandler.UpdateEmployeeDepartment)
	user.Patch("/rehire/:uuid", middlewares.RequirePermission("user.rehire"), userHandler.RehireUser)

	user.Get("/export", middlewares.RequirePermission("user.export"), userHandler.ExportUsersTemplateToExcel)
	user.Post("/import", middlewares.RequirePermission("user.import"), userHandler.ImportUsersFromExcel)
}
