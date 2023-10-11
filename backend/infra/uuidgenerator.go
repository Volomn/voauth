package infra

import "github.com/google/uuid"

type UUIDGenerator struct{}

func (gen *UUIDGenerator) New() (uuid.UUID, error) {
	return uuid.NewUUID()
}
