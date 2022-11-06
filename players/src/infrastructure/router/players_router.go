package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "players/src/infrastructure/restful/service/players"
)

type PlayersRouter struct {
	validator *validator.Validate
}

func (v *PlayersRouter) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewPlayersRouter(e *echo.Echo, playersService *service.PlayersService) *echo.Echo {
	e.Validator = &PlayersRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/players")
	r.POST("", playersService.CreatePlayer)
	r.GET("/:id", playersService.GetPlayer)
	return e
}
