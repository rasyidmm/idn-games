package repository

import "github.com/opentracing/opentracing-go"

type GamePlayRepository interface {
	StartGame(span opentracing.Span, in interface{}) (interface{}, error)
	EndGame(span opentracing.Span, in interface{}) (interface{}, error)
	ListGamePlayByPlayerId(span opentracing.Span, in interface{}) (interface{}, error)
	GetGamePlay(span opentracing.Span, in interface{}) (interface{}, error)
	PauseGame(span opentracing.Span, in interface{}) (interface{}, error)
}
