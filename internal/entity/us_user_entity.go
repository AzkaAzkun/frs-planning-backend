package entity

import (
	"frs-planning-backend/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsVerified  bool      `json:"is_verified" gorm:"default:false"`
	PhoneNumber string    `json:"phone_number"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	Role        UserRole  `json:"role"`
	Timestamp
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

// func (u *User) BeforeSave(tx *gorm.DB) error {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	var err error
// 	u.Password, err = utils.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
