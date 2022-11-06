package quests

import "github.com/labstack/echo/v4"

type QuestsInputPort interface {
	CreateQuest(echo.Context, interface{}) (interface{}, error)
	GetQuest(echo.Context, interface{}) (interface{}, error)
	ListQuests(echo.Context) (interface{}, error)
}

type QuestsOutputPort interface {
	CreateQuestsResponse(interface{}) (interface{}, error)
	GetQuestResponse(interface{}) (interface{}, error)
	ListQuestResponse(interface{}) (interface{}, error)
}
