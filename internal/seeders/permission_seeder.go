package seeders

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

var actionsByResource = map[string][]string{
	"permission":           {"view", "view-all", "create", "update", "delete"},
	"role":                 {"view", "view-uuid", "view-name", "view-all", "create", "update", "delete"},
	"branch":               {"view", "view-all", "create", "update", "delete"},
	"company":              {"view", "view-all", "view-companies", "create", "update", "delete"},
	"user_detail":          {"view", "create", "update", "delete"},
	"user":                 {"view", "view-all", "view-all-datatable", "create", "update", "delete", "rehire", "export", "import"},
	"department_group":     {"view", "view-all", "view-all-dt", "create", "update", "delete"},
	"department":           {"view", "view-all", "view-all-dt", "create", "update", "delete"},
	"shift":                {"view", "view-all", "create", "update", "delete"},
	"company_detail_shift": {"view", "view-all", "create", "update", "delete", "view-by-shift-id"},
	"user_family":          {"view", "view-all", "view-all-dt", "create", "update", "delete"},
}

var customPermissions = []models.Permission{
	{Name: "Sync Role Permissions", Label: "role.sync"},
}

func SeedPermissions(db *gorm.DB, actorID uint) error {
	title := cases.Title(language.English)

	for resource, actions := range actionsByResource {
		for _, action := range actions {
			name := fmt.Sprintf("%s %s", title.String(action), title.String(resource))
			label := fmt.Sprintf("%s.%s", resource, action)

			var existing models.Permission
			if err := db.Where("name = ?", name).First(&existing).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				p := models.Permission{
					UUID:      uuid.New().String(),
					Name:      name,
					Label:     label,
					CreatedBy: actorID,
					ModifyBy:  actorID,
				}
				if err := db.Create(&p).Error; err != nil {
					return fmt.Errorf("failed to seed permission %s: %w", name, err)
				}
				fmt.Println("Permission created:", name)
			}
		}
	}

	for _, cp := range customPermissions {
		var existing models.Permission
		if err := db.Where("name = ?", cp.Name).First(&existing).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			cp.UUID = uuid.New().String()
			cp.CreatedBy = actorID
			cp.ModifyBy = actorID

			if err := db.Create(&cp).Error; err != nil {
				return fmt.Errorf("failed to seed custom permission %s: %w", cp.Name, err)
			}
			fmt.Println("Custom permission created:", cp.Name)
		}
	}

	return nil
}
