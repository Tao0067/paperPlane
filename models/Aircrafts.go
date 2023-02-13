package models

import (
	"errors"
	global "go_xx/config"
	"go_xx/utils"
	"log"
	"time"
)

type Aircraft struct {
	Id           uint            `json:"id"`
	Name         string          `json:"name"`
	CoverPicture string          `json:"cover_picture"`
	Tag          string          `json:"tag"`
	Type         string          `json:"type"`
	Content      string          `json:"content"`
	CreateAd     utils.LocalTime `json:"create_ad"`
	UpdateAd     utils.LocalTime `json:"update_ad"`
}

func (a *Aircraft) CreateAircraft() error {
	a.CreateAd = utils.LocalTime(time.Now())
	a.UpdateAd = utils.LocalTime(time.Now())
	tx := global.DB.Create(a)
	if tx.Error != nil {
		log.Fatal(tx.Error.Error())
		return errors.New("创建失败")
	}
	return nil
}

func (a *Aircraft) UpdateAircraft() error {
	a.UpdateAd = utils.LocalTime(time.Now())
	tx := global.DB.Updates(a)
	if tx.Error != nil {
		log.Fatal(tx.Error.Error())
		return errors.New("更新失败")
	}
	return nil
}

func FindAircraftById(id uint) (aircraft *Aircraft, err error) {
	tx := global.DB.First(&aircraft, id)
	if tx.Error != nil {
		log.Fatal(tx.Error.Error())
		return aircraft, errors.New("更新失败")
	}
	return
}

func QueryPageAircraft() {

}
