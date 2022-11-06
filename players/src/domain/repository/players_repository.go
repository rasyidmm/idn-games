package repository

import "github.com/opentracing/opentracing-go"

type PlayersRepository interface {
	CreatePlayer(span opentracing.Span, in interface{}) (interface{}, error)
	GetPlayer(span opentracing.Span, in interface{}) (interface{}, error)
}
