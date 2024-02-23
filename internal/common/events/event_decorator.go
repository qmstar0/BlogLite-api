package events

import "github.com/qmstar0/eio"

type EventDecorator struct {
	Event any
}

func Wrap(e any) Serializable {
	return EventDecorator{
		Event: e,
	}
}

func (e EventDecorator) MarshalBinary() (data []byte, err error) {
	return eventCodec.Encode(e.Event)
}

func (e EventDecorator) UnmarshalBinary(data []byte) error {
	return eventCodec.Decode(data, &e.Event)
}

var eventCodec eio.Codec

func init() {
	eventCodec = eio.NewGobCodec()
}
