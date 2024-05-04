package eventbus

import (
	"context"
	"go-blog-ddd/internal/domain/domainevent"
	"sync"
)

type EventBus interface {
	Channel(ctx context.Context) <-chan domainevent.DomainEvent
	Publish(evts domainevent.DomainEvent)
}

func NewEventBus(eventBufferSize int) EventBus {
	return &eventBus{
		clients:         make([]*client, 0),
		clientsLock:     sync.RWMutex{},
		eventBufferSize: eventBufferSize,
	}
}

type eventBus struct {
	clients     []*client
	clientsLock sync.RWMutex
	//clientsWg   sync.WaitGroup

	eventBufferSize int

	//closed   bool
	//closing  chan struct{}
	//closLock sync.Mutex
}

func (e *eventBus) Channel(ctx context.Context) <-chan domainevent.DomainEvent {
	e.clientsLock.Lock()
	defer e.clientsLock.Unlock()

	cli := &client{
		Ch:      make(chan domainevent.DomainEvent, e.eventBufferSize),
		sending: sync.Mutex{},
	}

	go func(c *client, bus *eventBus) {
		<-ctx.Done()
		c.Close()

		bus.clientsLock.Lock()
		defer bus.clientsLock.Unlock()

		bus.removeClient(c)
	}(cli, e)

	e.addClient(cli)
	return cli.Ch
}

func (e *eventBus) Publish(evt domainevent.DomainEvent) {
	e.clientsLock.RLock()
	defer e.clientsLock.RUnlock()

	e.sendEvent(evt)
}

func (e *eventBus) sendEvent(evt domainevent.DomainEvent) {
	for _, c := range e.clients {
		go c.sendToChannel(evt)
	}
}

func (e *eventBus) addClient(cli *client) {
	e.clients = append(e.clients, cli)
}
func (e *eventBus) removeClient(cli *client) {
	for i, c := range e.clients {
		if c == cli {
			e.clients = append(e.clients[:i], e.clients[i+1:]...)
			break
		}
	}
}

type client struct {
	Ch      chan domainevent.DomainEvent
	sending sync.Mutex

	closed  bool
	closing chan struct{}
}

func (c *client) sendToChannel(evt domainevent.DomainEvent) {
	c.sending.Lock()
	defer c.sending.Unlock()

	select {
	case <-c.closing:
	case c.Ch <- evt:
	}
}

func (c *client) Close() {
	c.sending.Lock()
	defer c.sending.Unlock()

	close(c.Ch)
}
