package repository

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/model/converter"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(db *gorm.DB, request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error)
	// Get(db *gorm.DB, request *model.)
	// GetMany()
	// Update()
	// Delete()
}

type organizationRepository struct {
	Log *zap.Logger
}

func NewOrganizationRepository(log *zap.Logger) OrganizationRepository {
	return &organizationRepository{
		Log: log,
	}
}

func (repository *organizationRepository) Create(db *gorm.DB, request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error) {
	organization := entity.Organization{
		Name:        request.Name,
		Description: request.Description,
	}

	if result := db.Create(&organization); result.Error != nil {
		return nil, result.Error
	}

	response := converter.OrganizationToResponse(organization)
	return &response, nil
}

// func (repository *organizationRepository)
