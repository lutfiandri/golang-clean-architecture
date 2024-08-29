package model

import "time"

type BaseResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"udpated_at"`
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
	Data         any          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Error            string                    `json:"error"`
	ValidationErrors []ValidationErrorResponse `json:"validation_errors"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

// helper
func NewResponse(data any) Response {
	return Response{
		Data: data,
	}
}

func NewPageResponse[T any](data []T, metadata PageMetadata) PageResponse {
	return PageResponse{
		Data:         data,
		PageMetadata: metadata,
	}
}
func NewErrorResponse(err error, validationErrors []ValidationErrorResponse) ErrorResponse {
	return ErrorResponse{
		Error:            err.Error(),
		ValidationErrors: validationErrors,
	}
}