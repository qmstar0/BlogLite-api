package domainevent

import "github.com/google/uuid"

func NewEventID() string {
	return uuid.New().String()
}
