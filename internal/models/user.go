package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
	ID 				uuid.UUID 	`json:"id" gorm:"type:uuid;primaryKey"`
	FirstName 		string 		`json:"first_name"`
	LastName 		string 		`json:"last_name"`
	Email	 		string 		`json:"email" gorm:"index;unique;not null"`
	Phone	 		string 		`json:"phone"`
	Password 		string 		`json:"password"`
	Code 			int 		`json:"code"`
	ExpirationTime 	time.Time 	`json:"expiration_time"`
	IsVerified	 	bool 		`json:"is_verfied" gorm:"default:false"`
	UserType	 	string 		`json:"user_type" gorm:"default:buyer"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:current_timestamp"`
}

// BeforeCreate hook will be called automatically by GORM
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}