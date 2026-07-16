package infra

import "github.com/google/uuid"

type UuidProvider struct {
}

func (u UuidProvider) GetId() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
