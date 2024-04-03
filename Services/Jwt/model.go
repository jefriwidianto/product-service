package Jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"product-service/Config"
)

type JwtConfig struct {
	Config *Config.Jwt
}

var JwtConfigValue JwtConfig

type config struct {
	Auth *jwtauth.JWTAuth
}

type Jwt interface {
	Encode(ParamKeys jwt.MapClaims) (token string, err error)
	Decode(token string) (res map[string]interface{}, err error)
}

func AuthKey() Jwt {
	return &config{Auth: jwtauth.New(JwtConfigValue.Config.Encrypt, []byte(JwtConfigValue.Config.SecretKey), nil)}
}
