package repositories

type User struct {
	UserId      int    `json:"user_id"`
	FirstName   string `json:"first_name" db:"first_name"`
	SecondName  string `json:"second_name" db:"second_name"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password" db:"password"`
}

type UserSignIn struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
