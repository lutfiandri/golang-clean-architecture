package model

import (
	"time"
)

type PageRequest struct {
	Page *uint `query:"page"`
	Size *uint `query:"size"`
}

type BaseResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Type  string `json:"type"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Response struct {
	Data any `json:"data,omitempty"`
}

type PageResponse struct {
	Data     any       `json:"data,omitempty"`
	PageMeta *PageMeta `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Error            string                     `json:"error"`
	ValidationErrors []*ValidationErrorResponse `json:"validation_errors"`
}

type PageMeta struct {
	Page      uint   `json:"page"`
	Size      uint   `json:"size"`
	TotalItem uint64 `json:"total_item"`
	TotalPage uint64 `json:"total_page"`
}

// helper
func NewResponse(data any) *Response {
	return &Response{
		Data: data,
	}
}

func NewPageResponse(data any, metadata *PageMeta) *PageResponse {
	return &PageResponse{
		Data:     data,
		PageMeta: metadata,
	}
}

func NewErrorResponse(err error, validationErrors []*ValidationErrorResponse) *ErrorResponse {
	return &ErrorResponse{
		Error:            err.Error(),
		ValidationErrors: validationErrors,
	}
}
