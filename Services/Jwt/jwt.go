package Jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func (c config) Encode(ParamKeys jwt.MapClaims) (token string, err error) {
	_, token, err = c.Auth.Encode(ParamKeys)
	return
}

func (c config) Decode(token string) (res map[string]interface{}, err error) {
	jwtToken, err := c.Auth.Decode(token)
	if err != nil {
		return
	}
	res = jwtToken.PrivateClaims()
	return
}
