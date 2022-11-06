package players

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"players/src/shared/tracing"
	usecase "players/src/usecase/players"
)

type CreatePlayerRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
type GetPlayerRequest struct {
	Id string `json:"id"`
}

type PlayersService struct {
	uc usecase.PlayersInputPort
}

func NewPlayersService(u usecase.PlayersInputPort) *PlayersService {
	return &PlayersService{
		uc: u,
	}
}

func (s *PlayersService) CreatePlayer(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "CreatePlayer")
	defer sp.Finish()

	reqData := new(CreatePlayerRequest)
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

	var request *usecase.CreatePlayerRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.CreatePlayer(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}

func (s *PlayersService) GetPlayer(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "GetPlayer")
	defer sp.Finish()

	id := ctx.Param("id")
	reqData := &GetPlayerRequest{id}
	tracing.LogRequest(sp, reqData)

	err := ctx.Validate(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var request *usecase.GetPlayerRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.GetPlayer(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)

}
