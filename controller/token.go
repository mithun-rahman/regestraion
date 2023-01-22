package controller

import (
	"RegLog/model"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

func (c *Controller) GenerateToken(user *model.Users) (model.Token, error) {
	randStr := GenerateRandomString(32)

	// generate a access token
	accessToken, err := c.generateAccessToken(user, randStr)
	if err != nil {
		return model.Token{}, err
	}

	// generate a refresh token
	refreshToken, err := c.generateRefreshToken(user, randStr)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (c *Controller) generateAccessToken(tPayload *model.Users, randStr string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := jwt.MapClaims{
		"jti": randStr,
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Unix(),

		"email":     tPayload.Email,
		"user_name": tPayload.User_name,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(c.PrivateKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c *Controller) generateRefreshToken(tPayload *model.Users, randStr string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := jwt.MapClaims{
		"jti": randStr,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
		"iat": time.Now().Unix(),

		"email":     tPayload.Email,
		"user_name": tPayload.User_name,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(c.PrivateKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
