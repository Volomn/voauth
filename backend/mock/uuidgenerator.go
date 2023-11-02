package mock

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUUIDGenerator struct{ mock.Mock }

func (gen *MockUUIDGenerator) New() (uuid.UUID, error) {
	args := gen.Called()
	return args.Get(0).(uuid.UUID), args.Error(1)
}
