package converter

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		BaseResponse:  *BaseEntityToResponse(&user.BaseEntity),
		Name:          user.Name,
		Email:         user.Email,
		Role:          RoleToResponse(&user.Role),
		Organizations: OrganizationToResponseMany(user.Organizations),
	}
}
