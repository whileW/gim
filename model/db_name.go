package model

import (
	"github.com/jinzhu/gorm"
	"github.com/whileW/enze-global"
)

const ImDB =	"imdb"

func GetDb() *gorm.DB {
	return global.GVA_DB.Get(ImDB)
}