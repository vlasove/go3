package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/vlasove/userserver/proto/user"
)

var key = []byte("MyNewSecretKey99999")

type MyClaim struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Encode(user *pb.User) (string, error)
	Decode(token string) (*MyClaim, error)
}

type TokenService struct {
	repo Repository
}

func (ts *TokenService) Decode(token string) (interface{}, error) {
	tokenType, err := jwt.ParseWithClaims(string(key), &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claim, isOk := tokenType.Claims.(*MyClaim); isOk && tokenType.Valid {
		return claim, nil
	}
	return nil, err

}

func (ts *TokenService) Encode(user *pb.User) (string, error) {
	claim := MyClaim{
		user,
		jwt.StandardClaims{
			ExpiresAt: 20000,
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(key)
}
