package router

import (
	service "game-play/src/infrastructure/restful/service/game_play"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GamePlayRouter struct {
	validator *validator.Validate
}

func (v *GamePlayRouter) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewGamePlayRouter(e *echo.Echo, gamePlayService *service.GamePlayService) *echo.Echo {
	e.Validator = &GamePlayRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/game-play")
	r.POST("/start-game", gamePlayService.StartGame)
	r.POST("/end-game", gamePlayService.EndGame)
	r.POST("/pause-game", gamePlayService.PauseGame)
	r.POST("/list-by-player-id", gamePlayService.ListGamePlayByPlayerId)
	r.GET("/:id", gamePlayService.GetGamePlay)
	return e
}
