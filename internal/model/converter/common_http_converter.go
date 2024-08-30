package converter

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
)

func BaseEntityToResponse(baseEntity *entity.BaseEntity) *model.BaseResponse {
	return &model.BaseResponse{
		ID:        baseEntity.ID,
		CreatedAt: baseEntity.CreatedAt,
		UpdatedAt: baseEntity.UpdatedAt,
	}
}
