package controllers

type CreatNewsPostRequest struct {
	CoverImage string `json:"cover_image" binding:"required"`
	Title      string `json:"title" binding:"required"`
	SubTitle   string `json:"sub_title" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type CreatNewsPostResponse struct {
	Message string `json:"message"`
}
