package usecase

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/delivery/http/exception"
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/model/converter"
	"github.com/lutfiandri/golang-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type OrganizationUseCase interface {
	Create(request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error)
	GetMany(request *model.GetManyOrganizationRequest) ([]*model.OrganizationResponse, *model.PageMeta, error)
	Get(request *model.GetOrganizationRequest) (*model.OrganizationResponse, error)
	Update(request *model.UpdateOrganizationRequest) (*model.OrganizationResponse, error)
	Delete(request *model.DeleteOrganizationRequest) error
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

	organization := entity.Organization{
		Name:        request.Name,
		Description: request.Description,
	}

	err := usecase.organizationRepository.Create(tx, &organization)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := converter.OrganizationToResponse(&organization)
	return response, nil
}

func (usecase *organizationUseCase) GetMany(request *model.GetManyOrganizationRequest) ([]*model.OrganizationResponse, *model.PageMeta, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	organizations, pageMeta, err := usecase.organizationRepository.GetMany(tx, &request.PageRequest)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	response := converter.OrganizationToResponseMany(organizations)
	return response, pageMeta, nil
}

func (usecase *organizationUseCase) Get(request *model.GetOrganizationRequest) (*model.OrganizationResponse, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, err := usecase.organizationRepository.Get(tx, &request.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := converter.OrganizationToResponse(result)
	return response, nil
}

func (usecase *organizationUseCase) Update(request *model.UpdateOrganizationRequest) (*model.OrganizationResponse, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	if org, err := usecase.organizationRepository.Get(tx, &request.ID); org == nil || err != nil {
		tx.Rollback()
		return nil, exception.ErrOrganizationNotFound
	}

	organization := entity.Organization{
		BaseEntity:  entity.BaseEntity{ID: request.ID},
		Name:        request.Name,
		Description: request.Description,
	}

	err := usecase.organizationRepository.Update(tx, &organization.ID, &organization)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := converter.OrganizationToResponse(&organization)
	return response, nil
}

func (usecase *organizationUseCase) Delete(request *model.DeleteOrganizationRequest) error {
	tx := usecase.db.Begin()
	defer tx.Commit()

	if org, err := usecase.organizationRepository.Get(tx, &request.ID); org == nil || err != nil {
		tx.Rollback()
		return exception.ErrOrganizationNotFound
	}

	err := usecase.organizationRepository.Delete(tx, &request.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
