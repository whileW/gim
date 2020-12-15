package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/whileW/enze-global"
	"gim/model/request"
	"time"
)

//生成员工token
func GenerateTokenUser(id,name,head_img string) (string,error) {
	claims := request.CustomClaims{
		ID:          id,
		NickName:    name,
		HeadImg:	head_img,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       		// 签名生效时间
			Issuer:    "geeran-gim",                       // 签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	SigningKey := global.GVA_CONFIG.Setting.GetStringd("jwt_key","user_signing_key")
	token_str,err := token.SignedString([]byte(SigningKey))
	return token_str,err
}
//解析员工token
func ParseTokenUser(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		SigningKey := global.GVA_CONFIG.Setting.GetStringd("jwt_key","user_signing_key")
		return []byte(SigningKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("TokenMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("TokenExpired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("TokenNotValidYet")
			} else {
				return nil, errors.New("TokenInvalid")
			}
		}
	}
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("TokenInvalid")
}