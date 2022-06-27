package models

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Credential struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type JwtToken struct {
	Token    string
	ExpireAt time.Time
}
