package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hris_backend/database"
	"hris_backend/pkg"
	"hris_backend/routes/auth"
	"hris_backend/routes/branch"
	"hris_backend/routes/company"
	"hris_backend/routes/role"
	"hris_backend/routes/user"
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
}
