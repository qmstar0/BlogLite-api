package util

import "github.com/google/uuid"

func NewUUIDAutofill() string {
	return uuid.New().String() + "-autofill"
}
