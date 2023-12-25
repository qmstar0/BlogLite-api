package CQRS

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type JSONMarshal struct {
	GenIdFunc func() string
}

func (J *JSONMarshal) Marshal(value any) (*message.Message, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, ValueMarshalErr
	}
	return message.NewMessage(J.newId(), bytes), nil
}

func (J *JSONMarshal) Unmarshal(msg *message.Message, value any) error {
	if json.Unmarshal(msg.Payload, value) != nil {
		return ValueUnMarshalErr
	}
	return nil
}

func (J *JSONMarshal) newId() string {
	if J.GenIdFunc != nil {
		return J.GenIdFunc()
	}
	return watermill.NewUUID()
}
