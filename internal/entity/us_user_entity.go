package entity

import (
	"film-management-api-golang/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)

func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	Role        UserRole  `json:"role"`

	FilmLists []FilmList `json:"film_lists" gorm:"foreignKey:UserId"`
	Reviews   []Review   `json:"reviews" gorm:"foreignKey:UserId"`

	Timestamp
}

func (u *User) TableName() string {
	return "us_users"
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
