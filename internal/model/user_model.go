package model

type UserResponse struct {
	BaseResponse
	Role          RoleResponse           `json:"role"`
	Name          string                 `json:"name"`
	Email         string                 `json:"email"`
	Organizations []OrganizationResponse `json:"organizations"`
}

type GetUserRequest struct {
}
