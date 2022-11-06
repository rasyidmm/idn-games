package repository

import "github.com/opentracing/opentracing-go"

type QuestsRepository interface {
	CreateQuest(span opentracing.Span, in interface{}) (interface{}, error)
	ListQuests(span opentracing.Span) (interface{}, error)
	GetQuest(span opentracing.Span, in interface{}) (interface{}, error)
}
