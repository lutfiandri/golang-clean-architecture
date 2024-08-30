package repository

import (
	"log"

	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/helper"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/model/converter"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(db *gorm.DB, request *model.CreateOrganizationRequest) (*model.OrganizationResponse, error)
	GetMany(db *gorm.DB, request *model.GetManyOrganizationRequest) ([]*model.OrganizationResponse, *model.PageMetadata, error)
	Get(db *gorm.DB, request *model.GetOrganizationRequest) (*model.OrganizationResponse, error)
	Update(db *gorm.DB, request *model.UpdateOrganizationRequest) (*model.OrganizationResponse, error)
	Delete(db *gorm.DB, request *model.DeleteOrganizationRequest) error
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

	response := converter.OrganizationToResponse(&organization)
	return response, nil
}

func (repository *organizationRepository) GetMany(db *gorm.DB, request *model.GetManyOrganizationRequest) ([]*model.OrganizationResponse, *model.PageMetadata, error) {
	var organizations []*entity.Organization

	result := db.Scopes(helper.PaginateGorm(request.Page, request.Size)).Find(&organizations)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	response := converter.OrganizationToResponseMany(organizations)
	pageMetadata := helper.GetPageMetadata(db, &entity.Organization{}, request.Page, request.Size)

	return response, pageMetadata, nil
}

func (repository *organizationRepository) Get(db *gorm.DB, request *model.GetOrganizationRequest) (*model.OrganizationResponse, error) {
	var organization entity.Organization

	result := db.Where("id = ?", request.ID).First(&organization)
	if result.Error != nil {
		return nil, result.Error
	}

	response := converter.OrganizationToResponse(&organization)

	return response, nil
}

func (repository *organizationRepository) Update(db *gorm.DB, request *model.UpdateOrganizationRequest) (*model.OrganizationResponse, error) {
	organization := entity.Organization{
		BaseEntity:  entity.BaseEntity{ID: request.ID},
		Name:        request.Name,
		Description: request.Description,
	}

	log.Printf("%+v\n", organization)

	result := db.Save(&organization)
	if result.Error != nil {
		return nil, result.Error
	}

	response := converter.OrganizationToResponse(&organization)

	return response, nil
}

func (repository *organizationRepository) Delete(db *gorm.DB, request *model.DeleteOrganizationRequest) error {
	organization := entity.Organization{
		BaseEntity: entity.BaseEntity{ID: request.ID},
	}

	result := db.Delete(&organization)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
