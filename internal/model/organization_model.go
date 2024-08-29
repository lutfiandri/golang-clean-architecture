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
