package appy_http

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtProviderError struct {
	Message string
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type Claims = map[string]any

type JwtConfig struct {
	secret               string
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

func (j JwtProviderError) Error() string {
	return j.Message
}

func NewJwtConfig(secret string, accessTokenLifetime, refreshTokenLifetime time.Duration) JwtConfig {
	return JwtConfig{
		secret:               secret,
		accessTokenLifetime:  accessTokenLifetime,
		refreshTokenLifetime: refreshTokenLifetime,
	}
}

func (j JwtConfig) GenerateToken(lifetime time.Duration, sub any, claims Claims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	// Set claims
	tokenClaims := token.Claims.(jwt.MapClaims)
	tokenClaims["sub"] = sub
	tokenClaims["exp"] = time.Now().Add(lifetime).Unix()

	for key, value := range claims {
		tokenClaims[key] = value
	}

	// Sign tokens with secret
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JwtConfig) GeneratePair(sub any, claims Claims) (*TokenPair, error) {
	result := &TokenPair{}

	accessToken, err := j.GenerateToken(j.accessTokenLifetime, sub, claims)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.GenerateToken(j.refreshTokenLifetime, sub, claims)
	if err != nil {
		return nil, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}

func (j JwtConfig) ParseBearerToken(token string) (any, Claims, error) {
	// Since this is bearer token we need to parse the token out
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	} else {
		return nil, nil, JwtProviderError{Message: "token is not a bearer token format: 'Bearer <token>'"}
	}

	return j.ParseToken(token)
}

func (j JwtConfig) ParseToken(token string) (any, Claims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !jwtToken.Valid {
		return nil, nil, JwtProviderError{Message: "token invalid"}
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, JwtProviderError{Message: "token missing claims"}
	}

	jwtToken.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), true)

	// Get the subject
	sub, ok := claims["sub"]
	if !ok {
		return nil, nil, JwtProviderError{Message: "token missing sub"}
	}

	return sub, claims, nil
}
