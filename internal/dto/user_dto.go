package dto

type (
	RegisterRequest struct {
		Username    string `json:"username" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required"`
		DisplayName string `json:"display_name" binding:"required"`
		Bio         string `json:"bio" binding:""`
	}

	RegisterResponse struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
		Bio         string `json:"bio"`
	}

	UserResponse struct {
		ID          string             `json:"id"`
		Username    string             `json:"username"`
		DisplayName string             `json:"display_name"`
		Bio         string             `json:"bio"`
		FilmLists   []FilmListResponse `json:"film_lists"`
		Reviews     []ReviewResponse   `json:"reviews"`
	}
)
