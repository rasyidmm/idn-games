package game_play

import (
	"errors"
	"game-play/src/domain/entity"
	"game-play/src/domain/repository"
	"game-play/src/shared/enum"
	"game-play/src/shared/tracing"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type StartGameRequest struct {
	Id        string
	PlayerId  string
	QuestId   string
	Status    string
	LimitTime string
}
type StartGameResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	QuestStart string
}
type EndGameRequest struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type EndGameResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Score      string
	LimitTime  string
	Percentage string
}

type ListGamePlayByPlayerIdRequest struct {
	PlayerId string
}

type GamePlay struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type ListGamePlayByPlayerIdResponse struct {
	GamePlay []GamePlay
}

type GetGamePlayRequest struct {
	Id string
}
type GetGamePlayResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type PauseGameRequest struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type PauseGameResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Score      string
	LimitTime  string
	Percentage string
}

type GamePlayInteractor struct {
	repo repository.GamePlayRepository
	out  GamePlayOutputPort
}

func NewGamePlayInteractor(r repository.GamePlayRepository, o GamePlayOutputPort) *GamePlayInteractor {
	return &GamePlayInteractor{
		repo: r,
		out:  o,
	}
}

func (u *GamePlayInteractor) StartGame(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*StartGameRequest)
	var request *entity.StartGameRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.StartGame(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.StartGameResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.StartGameResponse(res)
}

func (u *GamePlayInteractor) EndGame(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*EndGameRequest)
	var request *entity.EndGameRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.EndGame(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.EndGameResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.EndGameResponse(res)
}

func (u *GamePlayInteractor) ListGamePlayByPlayerId(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*ListGamePlayByPlayerIdRequest)

	var (
		request *entity.ListGamePlayByPlayerIdRequest
		res     *entity.ListGamePlayByPlayerIdResponse
	)

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.ListGamePlayByPlayerId(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.ListGamePlayByPlayerIdResponse(res)
}

func (u *GamePlayInteractor) GetGamePlay(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*GetGamePlayRequest)

	var (
		request *entity.GetGamePlayRequest
		res     *entity.GetGamePlayResponse
	)

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.GetGamePlay(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.GetGamePlayResponse(res)
}

func (u *GamePlayInteractor) PauseGame(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*PauseGameRequest)
	var request *entity.PauseGameRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	resData, err := u.repo.PauseGame(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.PauseGameResponse
	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return u.out.EndGameResponse(res)
}
