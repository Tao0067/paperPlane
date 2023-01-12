package models

import (
	"errors"
	"go.uber.org/zap"
	global "go_xx/config"
	"time"
)

type AdminUser struct {
	Id        int
	Name      string
	Password  string
	CreateAd  time.Time
}

func (u *AdminUser) Create() error  {
	u.CreateAd = time.Now()
	tx := global.DB.Create(u)
	if tx.RowsAffected == 0 {
		zap.S().Info("管理员创建失败！", tx.Error)
		return errors.New(tx.Error.Error())
	}

	return nil
}

func AdminUserByNameAndPas(name string, password string) (adminUser AdminUser, err error)  {
	adminUser = AdminUser{}
	if tx := global.DB.Where("name = ? and password = ?", name, password).First(&adminUser);tx.RowsAffected == 0 {
		return
	}
	return
}

func FindAdminUserById(id int) (adminUser AdminUser, err error)  {
	adminUser = AdminUser{}
	if tx := global.DB.First(&adminUser, id); tx.Error != nil {
		return adminUser ,errors.New(tx.Error.Error())
	}
	return adminUser, nil
}
