package inmemory

import (
	"context"
	"zenrows-proxy/kit/query"
)

// QueryBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	queryHandlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		queryHandlers: make(map[query.Type]query.Handler),
	}
}

// Ask implements the bus.Bus interface.
func (b *QueryBus) Ask(ctx context.Context, q query.Query) (query.Response, error) {
	handler, ok := b.queryHandlers[q.Type()]
	if !ok {
		return nil, nil
	}

	return handler.Handle(ctx, q)
}

// Register implements the bus.Bus interface.
func (b *QueryBus) Register(q query.Type, handler query.Handler) {
	b.queryHandlers[q] = handler
}
