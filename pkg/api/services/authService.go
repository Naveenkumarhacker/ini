package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"ini/pkg/api/models"
	"time"
)

var AuthService authServiceI = new(authService)

type authServiceI interface {
	GenerateJWT(credential models.Credential) (*models.JwtToken, error)
	RefeshJWT(tokenString string) (*models.Claims, error)
}

type authService struct {
}

var JwtKey = []byte("secret_key")

func (a authService) GenerateJWT(credential models.Credential) (*models.JwtToken, error) {
	user, err := UserService.GetUserByUsernamAndPassword(credential.Name, credential.Password)
	if err != nil {
		return nil, err
	}

	if &user == nil {
		return nil, errors.New("UserName & Password not matched")
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &models.Claims{
		Name: credential.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return nil, err
	}

	return &models.JwtToken{
		Token:    tokenString,
		ExpireAt: expirationTime,
	}, nil
}

func (a authService) RefeshJWT(tokenString string) (*models.Claims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid { // Verification token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
