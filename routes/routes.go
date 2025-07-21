package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hris_backend/database"
	"hris_backend/pkg"
	"hris_backend/pkg/r2"
	"hris_backend/routes/auth"
	"hris_backend/routes/branch"
	"hris_backend/routes/company"
	"hris_backend/routes/department"
	departmentgroup "hris_backend/routes/department_group"
	"hris_backend/routes/employment_history"
	"hris_backend/routes/role"
	"hris_backend/routes/upload"
	"hris_backend/routes/user"
	usereducation "hris_backend/routes/user_education"
	"log"
	"os"
	"time"
)

func SetupRoutes(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
		MaxAge:           int((24 * time.Hour).Seconds()),
	}))

	r2Client, err := r2.NewR2Client(
		os.Getenv("R2_ENDPOINT"),
		os.Getenv("R2_ACCESS_KEY_ID"),
		os.Getenv("R2_SECRET_ACCESS_KEY"),
		true,
	)
	if err != nil {
		log.Fatalf("Failed to initialize R2 client: %v", err)
	}

	db := database.DB
	pkg.InitPermissionRepo(db)
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth.SetupAuthRoutes(v1, db)
	role.SetupRoleRoutes(v1, db)
	role.SetupRolePermissionRoutes(v1, db)
	role.SetupPermissionRoutes(v1, db)
	branch.SetupBranchRoutes(v1, db)
	company.SetupCompanyRoutes(v1, db)
	user.SetupUserDetailRoutes(v1, db)
	user.SetupUserRoutes(v1, db)
	departmentgroup.SetupDepartmentGroupRoutes(v1, db)
	department.SetupDepartmentRoutes(v1, db)
	employment_history.SetupEmploymentHistoryRoutes(v1, db)
	usereducation.SetupUserEducationRoutes(v1, db)
	upload.SetupUploadRoutes(v1, r2Client, os.Getenv("R2_BUCKET"))
}
