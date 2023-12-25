package CQRS

import (
	"errors"
	"github.com/ThreeDotsLabs/watermill/message"
)

var (
	ValueMarshalErr   = errors.New("序列化时发生错误")
	ValueUnMarshalErr = errors.New("反序列化时发生错误")
)

type CommandEventMarshal interface {
	Marshal(value any) (*message.Message, error)
	Unmarshal(msg *message.Message, value any) error
	Name(value any) string
}
