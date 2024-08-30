package repository

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/helper"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(db *gorm.DB, organization *entity.Organization) error
	GetMany(db *gorm.DB, pageMeta *model.PageRequest) ([]*entity.Organization, *model.PageMeta, error)
	Get(db *gorm.DB, id *uint) (*entity.Organization, error)
	Update(db *gorm.DB, id *uint, organization *entity.Organization) error
	Delete(db *gorm.DB, id *uint) error
}

type organizationRepository struct {
	Log *zap.Logger
}

func NewOrganizationRepository(log *zap.Logger) OrganizationRepository {
	return &organizationRepository{
		Log: log,
	}
}

func (repository *organizationRepository) Create(db *gorm.DB, organization *entity.Organization) error {
	if result := db.Create(&organization); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *organizationRepository) GetMany(db *gorm.DB, pageRequest *model.PageRequest) ([]*entity.Organization, *model.PageMeta, error) {
	var organizations []*entity.Organization

	result := db.Scopes(helper.PaginateGorm(pageRequest.Page, pageRequest.Size)).Find(&organizations)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	pageMeta, err := helper.GetPageMeta(db, &entity.Organization{}, pageRequest.Page, pageRequest.Size)
	if err != nil {
		return nil, nil, err
	}

	return organizations, pageMeta, nil
}

func (repository *organizationRepository) Get(db *gorm.DB, id *uint) (*entity.Organization, error) {
	var organization entity.Organization

	result := db.Where("id = ?", id).First(&organization)
	if result.Error != nil {
		return nil, result.Error
	}

	return &organization, nil
}

func (repository *organizationRepository) Update(db *gorm.DB, id *uint, organization *entity.Organization) error {
	result := db.Save(&organization)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *organizationRepository) Delete(db *gorm.DB, id *uint) error {
	organization := entity.Organization{
		BaseEntity: entity.BaseEntity{ID: *id},
	}

	result := db.Delete(&organization)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
