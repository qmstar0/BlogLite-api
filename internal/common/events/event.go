package events

import (
	"encoding"
	"github.com/qmstar0/eio"
	"time"
)

type Event struct {
	EventID     string
	AggregateID int
	Type        EventType
	Data        Serializable
	timestamp   time.Time
}

type Serializable interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

var EventCodec eio.Codec

func init() {
	EventCodec = eio.NewGobCodec()
}
