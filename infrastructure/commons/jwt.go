package commons

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	jwt.Token
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(id int, email string) (string, error) {
	claims := JWTClaims{
		jwt.Token{
			Claims: jwt.MapClaims{
				"id":  id,
				"exp": time.Now().Local().Add(time.Hour * time.Duration(jwtConf.ExpiresDuration)).Unix(),
				"iat": time.Now().Local().Unix(),
				"iss": "auth.service",
				"sub": email,
			},
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.Token.Claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}

func GetUserId(c echo.Context) int {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return int(claims["id"].(float64))
	}
	return 0
}
