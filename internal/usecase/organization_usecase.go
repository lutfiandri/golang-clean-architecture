package usecase

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type OrganizationUseCase interface {
	Create(request *model.CreateOrganizationRequest) (*model.Response, error)
	GetMany(request *model.GetManyOrganizationRequest) (*model.PageResponse, error)
	Get(request *model.GetOrganizationRequest) (*model.Response, error)
	Update(request *model.UpdateOrganizationRequest) (*model.Response, error)
	Delete(request *model.DeleteOrganizationRequest) (*model.Response, error)
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

func (usecase *organizationUseCase) Create(request *model.CreateOrganizationRequest) (*model.Response, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, err := usecase.organizationRepository.Create(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := model.NewResponse(result)
	return response, nil
}

func (usecase *organizationUseCase) GetMany(request *model.GetManyOrganizationRequest) (*model.PageResponse, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, pageMeta, err := usecase.organizationRepository.GetMany(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := model.NewPageResponse(result, pageMeta)
	return response, nil
}

func (usecase *organizationUseCase) Get(request *model.GetOrganizationRequest) (*model.Response, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, err := usecase.organizationRepository.Get(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := model.NewResponse(result)
	return response, nil
}

func (usecase *organizationUseCase) Update(request *model.UpdateOrganizationRequest) (*model.Response, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	result, err := usecase.organizationRepository.Update(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := model.NewResponse(result)
	return response, nil
}

func (usecase *organizationUseCase) Delete(request *model.DeleteOrganizationRequest) (*model.Response, error) {
	tx := usecase.db.Begin()
	defer tx.Commit()

	_, err := usecase.organizationRepository.Get(tx, &model.GetOrganizationRequest{ID: request.ID})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = usecase.organizationRepository.Delete(tx, request)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := model.NewResponse(nil)
	return response, nil
}
