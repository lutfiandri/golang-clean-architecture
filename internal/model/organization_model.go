package model

type OrganizationResponse struct {
	BaseResponse
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description"`
}
