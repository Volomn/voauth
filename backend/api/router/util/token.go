package util

import (
	"time"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/go-chi/jwtauth"
)

func CreateUserAccessToken(secretKey string, user domain.User) (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"authUUID": user.UUID.String(), "exp": time.Now().UTC().Add(24 * time.Hour)})
	return tokenString, err
}
