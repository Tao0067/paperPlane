package models

import (
	"errors"
	"go.uber.org/zap"
	global "go_xx/config"
	"go_xx/utils"
	"time"
)

type AdminUser struct {
	Id       uint            `json:"id"`
	Name     string          `json:"name"`
	Password string          `json:"password"`
	CreateAd utils.LocalTime `json:"create_ad"`
}

// Create /** 创建管理员
func (u *AdminUser) Create() error {
	u.CreateAd = utils.LocalTime(time.Now())
	u.Password = Encrypt(u.Password)
	tx := global.DB.Create(u)
	if tx.RowsAffected == 0 {
		zap.S().Info("管理员创建失败！", tx.Error)
		return errors.New(tx.Error.Error())
	}

	return nil
}

func FindAdminUserByNameAndPas(name string, password string) (adminUser AdminUser, err error) {
	adminUser = AdminUser{}
	password = Encrypt(password)
	if tx := global.DB.Where("name = ? and password = ?", name, password).First(&adminUser); tx.RowsAffected == 0 {
		return
	}
	return
}

func FindAdminUserById(id int) (adminUser AdminUser, err error) {
	adminUser = AdminUser{}
	if tx := global.DB.First(&adminUser, id); tx.Error != nil {
		return adminUser, errors.New(tx.Error.Error())
	}
	return adminUser, nil
}
