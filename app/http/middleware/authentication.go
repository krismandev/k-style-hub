package middleware

import (
	"k-style-test/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	UserID     int `json:"user_id"`
	CustomerID int `json:"customemr_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, customerID int) (string, string, error) {
	claims := &JWTCustomClaims{
		userID,
		customerID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	expiredAtStr := jwt.NewNumericDate(time.Now().Add(time.Hour * 72)).Format("2006-01-02 15:04:05")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(config.Config.JwtSecret))

	return jwtToken, expiredAtStr, err
}

func GetAuthUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)

	return claims
}

func JWTAuth() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(config.Config.JwtSecret),
	}

	return echojwt.WithConfig(config)
	// return middleware.JWTWithConfig(config)
}
