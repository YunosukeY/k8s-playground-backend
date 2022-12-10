package app

type TodoForPostRequest struct {
	Content string `json:"content" binding:"required"`
}

type TodoForResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Mail struct {
	To  string  `json:"to" binding:"required,email"`
	Sub *string `json:"sub" binding:"required"`
	Msg *string `json:"msg" binding:"required"`
}
