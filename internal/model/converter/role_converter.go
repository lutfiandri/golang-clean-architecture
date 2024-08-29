package converter

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
)

func RoleToResponse(role entity.Role) model.RoleResponse {
	return model.RoleResponse{
		BaseResponse: BaseEntityToResponse(role.BaseEntity),
		Name:         role.Name,
	}
}
