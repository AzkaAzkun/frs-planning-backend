package dto

type (
	AddCollaboratorRequest struct {
		Email       string `json:"email" binding:"required"`
		Workspaceid string `json:"workspaceid" binding:"required"`
		IsVerified  bool   `json:"is_verified" binding:"required"`
		Permission  string `json:"permission" binding:"required"`
	}

	DeleteCollaboratorRequest struct {
		Email       string `json:"email"`
		Workspaceid string `json:"workspaceid" binding:"required"`
	}

	GetAllCollaboratorResponse struct {
		Email      string `json:"email"`
		Name       string `json:"name"`
		IsVerified bool   `json:"is_verified"`
		Permission string `json:"permission"`
	}

	CollaboratorResponse struct {
		UserId      string `json:"userid"`
		Workspaceid string `json:"workspaceid"`
		IsVerified  bool   `json:"is_verified"`
		Permission  string `json:"permission"`
	}
)
