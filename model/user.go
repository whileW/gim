package model

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/whileW/enze-global/utils"
)

//用户
type User struct {
	UUID 			string				`json:"uuid"`

	Name 			string				`json:"name"`
	Phone 			string				`json:"phone"`
	Password 		string				`json:"password"`
	Salt 			string				`json:"salt"`
	HeadImg			string				`json:"head_img"`

	utils.BaseModel
}

func (u *User)BeforeCreate(scope *gorm.Scope) error{
	if u.UUID != "" {
		scope.SetColumn("id", uuid.New().String())
	}
	if u.HeadImg != "" {
		scope.SetColumn("head_img", "https://dss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=3531671336,3780835954&fm=26&gp=0.jpg")
	}
	return nil
}
func (u *User)Login(LoginName,Password string) error {
	db := GetDb()
	if err := db.Model(u).First(u,"phone = ?",LoginName).Error;err != nil{
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("账号密码错误请重试")
		}
		return err
	}
	pwd := utils.MD5V([]byte(Password+u.Salt))
	if u.Password != pwd {
		return errors.New("账号密码错误请重试")
	}
	return nil
}
func (u *User)Reg() error {
	u.Salt = utils.RandomString(8)
	u.Password = utils.MD5V([]byte(u.Password+u.Salt))
	db := GetDb()
	return db.Create(u).Error
}