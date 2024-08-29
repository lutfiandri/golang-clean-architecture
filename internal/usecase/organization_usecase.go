package usecase

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type OrganizationUseCase interface {
	Create(request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error)
	// Get()
	// GetMany()
	// Update()
	// Delete()
}

type organizationUseCase struct {
	db                     *gorm.DB
	organizationRepository repository.OrganizationRepository
}

func NewOrganizationUseCase(db *gorm.DB, organizationRepository repository.OrganizationRepository) OrganizationUseCase {
	return &organizationUseCase{
		db:                     db,
		organizationRepository: organizationRepository,
	}
}

func (usecase *organizationUseCase) Create(request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, err := usecase.organizationRepository.Create(usecase.db, request)
	return result, err
}
