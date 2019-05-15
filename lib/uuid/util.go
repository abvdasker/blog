package uuid

import (
	"github.com/gofrs/uuid"
)

func New() uuid.UUID {
	newUUID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return newUUID
}
