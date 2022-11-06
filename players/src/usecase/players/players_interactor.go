package players

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"players/src/domain/entity"
	"players/src/domain/repository"
	"players/src/shared/enum"
	"players/src/shared/tracing"
)

type CreatePlayerRequest struct {
	Username string
	Password string
	Nickname string
	Email    string
}
type CreatePlayerResponse struct {
	StatusCode string
	StatusDesc string
}
type GetPlayerRequest struct {
	Id string
}
type GetPlayerResponse struct {
	Id       string
	Username string
	Nickname string
	Email    string
}
type PlayersInteractor struct {
	repo repository.PlayersRepository
	out  PlayerOutputPort
}

func NewPlayersInteractor(r repository.PlayersRepository, o PlayerOutputPort) *PlayersInteractor {
	return &PlayersInteractor{
		repo: r,
		out:  o,
	}
}

func (u *PlayersInteractor) CreatePlayer(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*CreatePlayerRequest)
	var request *entity.CreatePlayerRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.CreatePlayer(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.CreatePlayerResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.CreatePlayerResponse(res)
}

func (u *PlayersInteractor) GetPlayer(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*GetPlayerRequest)
	var request *entity.GetPlayerRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.GetPlayer(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res GetPlayerResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.GetPlayerResponse(res)
}
