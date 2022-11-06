package repository

import (
	"github.com/opentracing/opentracing-go"
)

type QuestRepository interface {
	CreateQuest(span opentracing.Span, in interface{}) (interface{}, error)
	CreateTask(span opentracing.Span, in interface{}) (interface{}, error)
	GetQuest(span opentracing.Span, in interface{}) (interface{}, error)
	ListQuest(span opentracing.Span) (interface{}, error)
	GetTask(span opentracing.Span, in interface{}) (interface{}, error)
	ListTask(span opentracing.Span) (interface{}, error)
	ListTaskByQuestId(span opentracing.Span, in interface{}) (interface{}, error)
}
