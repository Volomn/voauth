package app

import "github.com/Volomn/voauth/backend/domain"

type Application interface {
	SignupUser(firstName, lastName, email, password string) (domain.User, error)
	AuthenticateWithEmailAndPassword(email, password string) (domain.User, error)
	GetAuthSecretKey() string
}
