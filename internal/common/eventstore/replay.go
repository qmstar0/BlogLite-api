package eventstore

import "common/domainevent"

type Replayer[E any] map[uint16]func(raw any, entity *E) error

func (r Replayer[E]) Replay(events []domainevent.DomainEvent, entity *E) error {
	for _, event := range events {
		if f, ok := r[event.EventType]; ok {
			if err := f(event.Event, entity); err != nil {
				return err
			}
		}
	}
	return nil
}
