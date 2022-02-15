package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/inhun/crypto-trading-bot/models"
)

func hasParamHS256Token(secretKey string, claims models.ParamClaims) (*models.Token, error) {
	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &models.Token{SignedToken: signedToken, Type: "Bearer"}, nil
}

func nonParamHS256Token(secretKey string, claims models.NonParamClaims) (*models.Token, error) {
	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &models.Token{SignedToken: signedToken, Type: "Bearer"}, nil
}
