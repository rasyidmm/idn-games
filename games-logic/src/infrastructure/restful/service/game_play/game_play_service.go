package game_play

import (
	"games-logic/src/shared/tracing"
	usecase "games-logic/src/usecase/game_play"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type StartGameRequest struct {
	Id       string `json:"id"`
	PlayerId string `json:"player_id"`
	QuestId  string `json:"quest_id"`
}
type EndGameRequest struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
}
type ListGamePlayByPlayerIdRequest struct {
	PlayerId string `json:"player_id"`
}

type PauseGameRequest struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
}
type GamePlayService struct {
	uc usecase.GamePlayInputPort
}

func NewGamePlayService(u usecase.GamePlayInputPort) *GamePlayService {
	return &GamePlayService{
		uc: u,
	}
}

func (s *GamePlayService) StartGame(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "StartGame")
	defer sp.Finish()

	reqData := new(StartGameRequest)
	err := ctx.Bind(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogRequest(sp, reqData)

	err = ctx.Validate(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var request *usecase.StartGameRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.StartGame(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}

func (s *GamePlayService) EndGame(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "EndGame")
	defer sp.Finish()

	reqData := new(EndGameRequest)
	err := ctx.Bind(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogRequest(sp, reqData)

	err = ctx.Validate(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var request *usecase.EndGameRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.EndGame(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}

func (s *GamePlayService) ListGamePlayByPlayerId(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "ListGamePlayByPlayerId")
	defer sp.Finish()

	reqData := new(ListGamePlayByPlayerIdRequest)
	err := ctx.Bind(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogRequest(sp, reqData)

	err = ctx.Validate(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var request *usecase.ListGamePlayByPlayerIdRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.ListGamePlayByPlayerId(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}
func (s *GamePlayService) PauseGame(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "PauseGame")
	defer sp.Finish()

	reqData := new(PauseGameRequest)
	err := ctx.Bind(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogRequest(sp, reqData)

	err = ctx.Validate(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var request *usecase.PauseGameRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.PauseGame(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}
