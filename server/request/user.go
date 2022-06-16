package request

type User struct {
	UserId      int    `json:"user_id"`
	FirstName   string `json:"first_name" binding:"required"`
	SecondName  string `json:"second_name" binding:"required"`
	Patronymic  string `json:"patronymic" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UserSignIn struct {
	Id          int    `json:"id" db:"id"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
