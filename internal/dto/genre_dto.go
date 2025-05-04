package dto

import "film-management-api-golang/internal/pkg/meta"

type (
	GenreRequest struct {
		Name string `json:"name"`
	}

	GenreResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	GenreResponsePaginated struct {
		Data []GenreResponse
		Meta meta.Meta
	}
)
