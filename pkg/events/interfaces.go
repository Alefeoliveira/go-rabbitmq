package events

import (
	"sync"
	"time"
)

// evento
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

// executa a operação
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// gerenciador que dispara o evento
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dipatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
