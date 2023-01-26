package repo

import (
	"RegLog/app/response"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

func (p *PostresBrand) GenerateToken(user *response.Users) (response.Token, error) {
	randStr := GenerateRandomString(32)

	// generate a access token
	accessToken, err := p.generateAccessToken(user, randStr)
	if err != nil {
		return response.Token{}, err
	}

	// generate a refresh token
	refreshToken, err := p.generateRefreshToken(user, randStr)
	if err != nil {
		return response.Token{}, err
	}

	return response.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (p *PostresBrand) generateAccessToken(tPayload *response.Users, randStr string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := jwt.MapClaims{
		"jti": randStr,
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Unix(),

		"email":     tPayload.Email,
		"user_name": tPayload.User_name,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(p.PrivateKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (p *PostresBrand) generateRefreshToken(tPayload *response.Users, randStr string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := jwt.MapClaims{
		"jti": randStr,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
		"iat": time.Now().Unix(),

		"email":     tPayload.Email,
		"user_name": tPayload.User_name,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(p.PrivateKey)

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
