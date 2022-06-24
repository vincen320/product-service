package authtoken

import "github.com/dgrijalva/jwt-go"

type JWTPayload struct {
	UserId int `json:"id,omitempty"`
	jwt.StandardClaims
}
