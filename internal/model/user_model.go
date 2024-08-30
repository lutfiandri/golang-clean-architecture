package model

type UserResponse struct {
	BaseResponse
	Name          string                  `json:"name"`
	Email         string                  `json:"email"`
	Role          *RoleResponse           `json:"role"`
	Organizations []*OrganizationResponse `json:"organizations"`
}

type GetUserRequest struct {
}
