package repositories

type User struct {
	User_id  int    `json:"user_id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignIn struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
