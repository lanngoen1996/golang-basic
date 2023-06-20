package users

import (
	"time"

	"gorm.io/gorm"
)

type UserInterface interface {
	Get() (*[]User, error)
	Create(payload *User) (*User, error)
	Update(payload *User) (*User, error)
	Find(id string) (*User, error)
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

func NewUser(DB *gorm.DB) UserInterface {
	return &UserModel{
		DB,
	}
}

func (u *UserModel) Get() (*[]User, error) {
	var users []User
	if res := u.DB.Model(&User{}).Find(&users); res.Error != nil {
		return nil, res.Error
	}
	return &users, nil
}

func (u *UserModel) Create(payload *User) (*User, error) {
	user := *payload
	if res := u.DB.Model(&User{}).Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *UserModel) Find(id string) (*User, error) {
	var user User
	if res := u.DB.Model(&User{}).First(&user, id); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *UserModel) Update(payload *User) (*User, error) {
	if res := u.DB.Save(&payload); res.Error != nil {
		return nil, res.Error
	}
	return payload, nil
}

func (u *UserModel) Delete(d *User) error {
	if res := u.DB.Delete(&d); res.Error != nil {
		return res.Error
	}
	return nil
}
