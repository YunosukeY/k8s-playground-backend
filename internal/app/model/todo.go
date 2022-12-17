package model

type TodoForPostRequest struct {
	Content string `json:"content" binding:"required"`
}

type TodoForResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
