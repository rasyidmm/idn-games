package players

import "github.com/labstack/echo/v4"

type PlayersInputPort interface {
	CreatePlayer(echo.Context, interface{}) (interface{}, error)
	GetPlayer(echo.Context, interface{}) (interface{}, error)
}

type PlayerOutputPort interface {
	CreatePlayerResponse(interface{}) (interface{}, error)
	GetPlayerResponse(interface{}) (interface{}, error)
}
