package game_play

import "github.com/labstack/echo/v4"

type GamePlayInputPort interface {
	StartGame(echo.Context, interface{}) (interface{}, error)
	EndGame(echo.Context, interface{}) (interface{}, error)
	ListGamePlayByPlayerId(echo.Context, interface{}) (interface{}, error)
	PauseGame(echo.Context, interface{}) (interface{}, error)
}

type GamePlayOutputPort interface {
	StartGameResponse(interface{}) (interface{}, error)
	EndGameResponse(interface{}) (interface{}, error)
	ListGamePlayByPlayerIdResponse(interface{}) (interface{}, error)
	PauseGameResponse(interface{}) (interface{}, error)
}
