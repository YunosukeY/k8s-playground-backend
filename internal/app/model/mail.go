package model

type Mail struct {
	To  string  `json:"to" binding:"required,email"`
	Sub *string `json:"sub" binding:"required"`
	Msg *string `json:"msg" binding:"required"`
}
