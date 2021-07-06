package models

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenModel struct {
	UserId uint
	jwt.StandardClaims
}

func (b *TokenModel) TableName() string {
	return "token"
}
