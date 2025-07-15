package departmentgroup

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupDepartmentGroupRoutes(router fiber.Router, db *gorm.DB) {
	companyRepo := repositories.NewCompanyRepository(db)
	userRepo := repositories.NewUserRepository(db)

	departmentGroupRepo := repositories.NewDepartmentGroupRepositories(db)
	departmentGroupService := services.NewDepartmentGroupService(departmentGroupRepo, companyRepo, userRepo)
	departmentGroupHandler := handlers.NewDepartmentGroupHandler(departmentGroupService)

	departmentGroup := router.Group("/department-groups")
	departmentGroup.Use(middlewares.JWTMiddleware)

	departmentGroup.Post("/create", middlewares.RequirePermission("department_group.create"), departmentGroupHandler.Create)
	departmentGroup.Get("/get-all", middlewares.RequirePermission("department_group.view-all"), departmentGroupHandler.GetAll)
	departmentGroup.Get("/get/:uuid", middlewares.RequirePermission("department_group.view"), departmentGroupHandler.GetByUUID)
	departmentGroup.Get("/get-all/dt", middlewares.RequirePermission("department_group.view-all-dt"), departmentGroupHandler.GetAllDataTable)
	departmentGroup.Patch("/update/:uuid", middlewares.RequirePermission("department_group.update"), departmentGroupHandler.Update)
	departmentGroup.Delete("/delete/:uuid", middlewares.RequirePermission("department_group.delete"), departmentGroupHandler.Delete)
}
