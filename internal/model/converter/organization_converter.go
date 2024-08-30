package converter

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
)

func OrganizationToResponse(organization *entity.Organization) *model.OrganizationResponse {
	return &model.OrganizationResponse{
		BaseResponse: *BaseEntityToResponse(&organization.BaseEntity),
		Name:         organization.Name,
		Description:  organization.Description,
	}
}

func OrganizationToResponseMany(organizations []*entity.Organization) []*model.OrganizationResponse {
	var responses []*model.OrganizationResponse

	for _, organization := range organizations {
		responses = append(responses, OrganizationToResponse(organization))
	}

	return responses
}
