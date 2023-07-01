package model

import (
	"time"

	"github.com/lanngoen1996/golang-basic/core"
	"gorm.io/gorm"
)

type UserInterface interface {
	Get(d *[]User) error
	Create(d *User) error
	Update(d *User) error
	Find(id string, d *User) error
	Delete(d *User) error
}

type User struct {
	ID        uint      `gorm:"column:id" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updateAt"`
}

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel() UserInterface {
	return &UserModel{
		DB: core.DB,
	}
}

func (u *UserModel) Get(d *[]User) error {
	if res := u.DB.Model(&User{}).Find(&d); res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserModel) Create(d *User) error {
	if res := u.DB.Model(&User{}).Create(&d); res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserModel) Find(id string, d *User) error {
	if res := u.DB.Model(&User{}).First(&d, id); res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserModel) Update(d *User) error {
	if res := u.DB.Save(&d); res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserModel) Delete(d *User) error {
	if res := u.DB.Delete(&d); res.Error != nil {
		return res.Error
	}
	return nil
}
