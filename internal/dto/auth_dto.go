package dto

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	LoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	GetMe struct {
		PersonalInfo PersonalInfo `json:"personal_info"`
	}

	PersonalInfo struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
		Bio         string `json:"bio"`
		Role        string `json:"role"`
	}
)
