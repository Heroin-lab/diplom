package request

type User struct {
	UserId   int    `json:"user_id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
