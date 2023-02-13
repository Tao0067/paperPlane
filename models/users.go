package models

import (
	global "go_xx/config"
	"go_xx/utils"
	"time"
)

type User struct {
	Id       uint            `json:"id"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Name     string          `json:"name"`
	Status   byte            `json:"status"`
	CreateAd utils.LocalTime `json:"create_ad"`
	UpdateAd utils.LocalTime `json:"update_ad"`
}

func (u *User) CreateUser() error {
	u.CreateAd = utils.LocalTime(time.Now())
	u.UpdateAd = utils.LocalTime(time.Now())
	u.Password = Encrypt(u.Password)
	u.Status = 1
	tx := global.DB.Create(u)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *User) UpdateUser() error {
	u.UpdateAd = utils.LocalTime(time.Now())
	tx := global.DB.Updates(u)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func FindUserByEmailAndPassword(email string, password string) (user User, err error) {
	password = Encrypt(password)
	if tx := global.DB.Where("email = ? and password = ?", email, password).First(&user); tx.Error != nil {
		return
	}
	return
}

func FindUserByNameOrEmail(name string, email string) (user User, err error) {
	if tx := global.DB.Where("name = ? or email = ?", name, email).First(&user); tx.Error != nil {
		return
	}
	return
}

func FindUserById(id uint) (user *User, err error) {
	if tx := global.DB.Find(&user, id); tx.Error != nil {
		return
	}
	return
}
