package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"strings"
)

type Claims struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
	jwt.StandardClaims
}

func GetTokenFromAuthorization(c echo.Context) (string, error) {
	if c.Request().Header.Get("Authorization") == "" {
		return "", fmt.Errorf("Token is required")
	}
	authorization := c.Request().Header.Get("Authorization")
	token := strings.Split(authorization, " ")[1]
	return token, nil
}

func GetDataToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, claims)
	if err != nil {
		return nil, err
	}

	if token.Claims == nil {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}
