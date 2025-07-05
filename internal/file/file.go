package file

import "github.com/google/uuid"

type file struct {
	ID   uuid.UUID
	Name string
	Data []byte
}
