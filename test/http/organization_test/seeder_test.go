package organization_test

import (
	"log"

	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

func seedOrganization(db *gorm.DB) {
	descriptions := []string{
		"description 1",
		"description 2",
		"description 3",
		"description 4",
		"description 5",
	}

	organizations := []entity.Organization{
		{
			Name:        "org 1",
			Description: &descriptions[0],
		},
		{
			Name:        "org 2",
			Description: &descriptions[1],
		},
		{
			Name:        "org 3",
			Description: &descriptions[2],
		},
		{
			Name:        "org 4",
			Description: &descriptions[3],
		},
		{
			Name:        "org 5",
			Description: &descriptions[4],
		},
	}

	if err := db.Create(&organizations).Error; err != nil {
		log.Fatalf("failed to seed organizations: %v", err)
	}
}
