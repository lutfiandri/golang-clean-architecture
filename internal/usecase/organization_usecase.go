package usecase

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type OrganizationUseCase interface {
	Create(request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error)
	// Get()
	GetMany(request *model.GetManyOrganizationRequest) ([]*model.OrganizationResponse, *model.PageMetadata, error)
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

	result, err := usecase.organizationRepository.Create(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return result, nil
}

func (usecase *organizationUseCase) GetMany(request *model.GetManyOrganizationRequest) ([]*model.OrganizationResponse, *model.PageMetadata, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, pageMeta, err := usecase.organizationRepository.GetMany(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return result, pageMeta, nil
}
