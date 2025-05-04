package dto

type (
	ReviewRequest struct {
		FilmId  string `json:"film_id" binding:"required,uuid"`
		Rating  int    `json:"rating" binding:"required,gte=1,lte=10"`
		Comment string `json:"comment" binding:"required"`
	}

	ReviewResponse struct {
		Film    string `json:"film"`
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}
)
