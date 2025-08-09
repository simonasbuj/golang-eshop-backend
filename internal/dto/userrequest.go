package dto

type UserSignIn struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type UserSignUp struct {
	UserSignIn
	Phone		string	`json:"phone"`
}