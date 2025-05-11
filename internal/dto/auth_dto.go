package dto

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	ForgotPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	ChangePasswordRequest struct {
		NewPassword string `json:"new_password"`
	}

	GetMe struct {
		PersonalInfo PersonalInfo `json:"personal_info"`
	}

	PersonalInfo struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}
)
