package mock

import "github.com/stretchr/testify/mock"

type MockPasswordHasher struct{ mock.Mock }

func (hasher *MockPasswordHasher) HashPassword(password string) (string, error) {
	args := hasher.Called(password)
	return args.String(0), args.Error(1)
}

func (hasher *MockPasswordHasher) IsPasswordMatch(password, hashedPassword string) bool {
	args := hasher.Called(password, hashedPassword)
	return args.Bool(0)
}
