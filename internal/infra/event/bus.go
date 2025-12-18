package event

import (
	"context"
	"reflect"
	"sync"
	"time"

	"github.com/motixo/goat-api/internal/domain/service"
)

type Handler func(context.Context, any) error

type InMemoryPublisher struct {
	mu       sync.RWMutex
	handlers map[reflect.Type][]Handler
	logger   service.Logger
	wg       sync.WaitGroup // Track active handlers for graceful shutdown
}

func NewInMemoryPublisher(logger service.Logger) *InMemoryPublisher {
	return &InMemoryPublisher{
		handlers: make(map[reflect.Type][]Handler),
		logger:   logger,
	}
}

func (p *InMemoryPublisher) RegisterHandler(eventType reflect.Type, handler Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.handlers[eventType] = append(p.handlers[eventType], handler)
}

func (p *InMemoryPublisher) Publish(ctx context.Context, event any) error {

	eventType := reflect.TypeOf(event)
	p.mu.RLock()
	handlers, exists := p.handlers[eventType]
	p.mu.RUnlock()

	if !exists {
		return nil
	}

	for _, handler := range handlers {
		p.wg.Add(1)
		go func(h Handler, e any, et reflect.Type) {
			defer p.wg.Done()
			defer func() {
				if r := recover(); r != nil {
					p.logger.Error("Event handler panicked", "event", eventType.String(), "panic", r)
				}
			}()
			bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := h(bgCtx, e); err != nil {
				p.logger.Error("Event handler failed", "event", eventType.String(), "error", err)
			}
		}(handler, event, eventType)
	}

	return nil
}

// Wait blocks until all background handlers are finished.
func (p *InMemoryPublisher) Wait() {
	p.wg.Wait()
}
