package model

type OrganizationResponse struct {
	BaseResponse
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description"`
}

type CreateOrganizationRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type GetOrganizationRequest struct {
	ID uint `param:"id" validate:"required"`
}

type GetManyOrganizationRequest struct {
	PageRequest
}

type UpdateOrganizationRequest struct {
	ID          uint    `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type DeleteOrganizationRequest struct {
	ID uint `param:"id" validate:"required"`
}
