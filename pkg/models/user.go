package models

import (
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    *string `json:"first_Name" validate:"required, min=1, max=100"`
	LastName     *string `json:"last_Name" validate:"required, min=1, max=100"`
	Password     *string `json:"password,omitempty"  validate:"required, min=6"`
	Email        *string `json:"email" validate:"required, email"`
	Phone        *string `json:"phone" validate:"required"`
	Token        *string `json:"token"`
	UserType     *string `json:"user_type,"  validate:"required, eq=ADMIN|eq=USER|eq=TEACHER|eq=STUDENT"`
	RefreshToken *string `json:"refresh_Token"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db

}
func FindEmailCount(user *User) (*User, *gorm.DB, int64) {
	var getEmail User
	var count int64
	db := db.Model(&User{}).Where("email=?", user.Email).Count(&count)
	return &getEmail, db, count
}
func FindPhoneCount(user *User) (*User, *gorm.DB, int64) {
	var getEmail User
	var count int64
	db := db.Model(&User{}).Where("phone=?", user.Phone).Count(&count)
	return &getEmail, db, count
}
