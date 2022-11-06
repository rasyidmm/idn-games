package router

import (
	service "games-logic/src/infrastructure/restful/service/quests_service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type QuestsRouter struct {
	validator *validator.Validate
}

func (v *QuestsRouter) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewQuestsRouter(e *echo.Echo, questsService *service.QuestsService) *echo.Echo {
	e.Validator = &QuestsRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/games")
	r.POST("", questsService.CreateQuest)
	r.GET("", questsService.ListQuests)
	r.GET("/:id", questsService.GetQuest)
	return e
}
