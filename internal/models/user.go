package models

import (
	"time"

	"github.com/google/uuid"
)


type User struct {
	ID 				uuid.UUID 	`json:"id"`
	FirstName 		string 		`json:"first_name"`
	LastName 		string 		`json:"last_name"`
	Email	 		string 		`json:"email"`
	Phone	 		string 		`json:"phone"`
	Password 		string 		`json:"password"`
	Code 			int 		`json:"code"`
	ExpirationTime 	time.Time 	`json:"expiration_time"`
	IsVerified	 	bool 		`json:"is_verfied"`
	UserType	 	string 		`json:"user_type"`
}
