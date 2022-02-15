package models

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type NonParamClaims struct {
	AccessKey string
	Nonce     uuid.UUID
	jwt.StandardClaims
}

type ParamClaims struct {
	AccessKey    string
	Nonce        uuid.UUID
	QueryHash    string
	QueryHashAlg string
	jwt.StandardClaims
}

type Token struct {
	SignedToken string
	Type        string
}
