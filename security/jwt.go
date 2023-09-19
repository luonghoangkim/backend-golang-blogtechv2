package security

import (
	"backend-blogtechv2/model"
	"time"

	"github.com/golang-jwt/jwt"
)
const Secret_key = "hkjkhakhsgjhaskhdkajdhfkadjh"
func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaim{
		UserID: user.UserID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resutl , err := token.SignedString([]byte (Secret_key))

	if err != nil {
		return "", err
	}

	return resutl , nil
}