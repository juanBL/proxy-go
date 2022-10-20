package query

import "context"

type Bus interface {
	Ask(context.Context, Query) (Response, error)
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=querymocks --output=querymocks --name=Bus

type Type string

type Query interface {
	Type() Type
}

type Response interface{}

type Handler interface {
	Handle(context.Context, Query) (Response, error)
}
