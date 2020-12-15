package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID          string
	NickName    string
	HeadImg		string
	jwt.StandardClaims
}
