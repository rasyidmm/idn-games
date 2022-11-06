package game_play

import (
	"errors"
	"fmt"
	"games-logic/src/domain/entity"
	"games-logic/src/domain/repository"
	"games-logic/src/shared/enum"
	"games-logic/src/shared/tracing"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"strings"
	"time"
)

type StartGameRequest struct {
	Id       string
	PlayerId string
	QuestId  string
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
type PauseGameRequest struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
}
type PauseGameResponse struct {
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

type GamePlayInteractor struct {
	repo  repository.GamePlayRepository
	repog repository.QuestsRepository
	out   GamePlayOutputPort
}

func NewGamePlayInteractor(r repository.GamePlayRepository, rg repository.QuestsRepository, o GamePlayOutputPort) *GamePlayInteractor {
	return &GamePlayInteractor{
		repo:  r,
		repog: rg,
		out:   o,
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

	reqQuest := &entity.GetQuestRequest{Id: reqdata.QuestId}
	resQuest, err := u.repog.GetQuest(sp, reqQuest)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	outQuest := resQuest.(*entity.GetQuestResponse)

	request.LimitTime = outQuest.TimeQuest
	request.Status = "Start"
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

	reqGet := &entity.GetGamePlayRequest{reqdata.Id}
	getRes, err := u.repo.GetGamePlay(sp, reqGet)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	outGet := getRes.(*entity.GetGamePlayResponse)
	limit, _ := strconv.Atoi(outGet.LimitTime)
	start, _ := time.Parse("2006-01-02 15:04:05", outGet.QuestStart)

	if time.Now().Before(start.Add(time.Minute * time.Duration(limit))) {
		reqQuest := &entity.GetQuestRequest{outGet.QuestId}
		resQuest, errQ := u.repog.GetQuest(sp, reqQuest)
		if errQ != nil {
			tracing.LogError(sp, errQ)
			return nil, errQ
		}

		outQuest := resQuest.(*entity.GetQuestResponse)

		listTaksId := strings.Split(reqdata.ListTaskId, ";")
		countDone := len(listTaksId)
		countTask := len(outQuest.Tasks)
		if countTask != countDone {
			tracing.LogError(sp, errors.New("Ada task yang belum selesai"))
			return nil, errors.New("Ada task yang belum selesai")
		}
		var tempScore int
		for _, itemTask := range listTaksId {
			for _, task := range outQuest.Tasks {
				if itemTask == task.Id {
					score, _ := strconv.Atoi(task.ScoreTask)
					tempScore += score
				}
			}
		}
		Percentage := countDone * 100 / countTask
		request.Score = strconv.Itoa(tempScore)
		request.Status = "Done"
		request.LimitTime = "0"
		request.Percentage = strconv.Itoa(Percentage)
		request.QuestEnd = time.Now().Format("2006-01-02 15:04:05")
	} else {
		request.QuestEnd = time.Now().Format("2006-01-02 15:04:05")
		request.Status = "Failed"
		request.LimitTime = "0"
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

	reqGet := &entity.GetGamePlayRequest{reqdata.Id}
	getRes, err := u.repo.GetGamePlay(sp, reqGet)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	outGet := getRes.(*entity.GetGamePlayResponse)
	limit, _ := strconv.Atoi(outGet.LimitTime)
	start, _ := time.Parse("2006-01-02 15:04:05", outGet.QuestStart)
	Finish := start.Add(time.Minute * time.Duration(limit))

	if time.Now().Before(Finish) {
		reqQuest := &entity.GetQuestRequest{outGet.QuestId}
		resQuest, errQ := u.repog.GetQuest(sp, reqQuest)
		if errQ != nil {
			tracing.LogError(sp, errQ)
			return nil, errQ
		}

		outQuest := resQuest.(*entity.GetQuestResponse)

		listTaksId := strings.Split(reqdata.ListTaskId, ";")
		CountDone := len(listTaksId)
		CountTask := len(outQuest.Tasks)
		var tempScore int
		for _, itemTask := range listTaksId {
			for _, task := range outQuest.Tasks {
				if itemTask == task.Id {
					score, _ := strconv.Atoi(task.ScoreTask)
					tempScore += score
				}
			}
		}
		var Percentage = CountDone * 100 / CountTask
		Tiiii := time.Now()
		fmt.Println(Tiiii)
		diff := Finish.Sub(time.Now())
		request.Score = strconv.Itoa(tempScore)
		request.Status = "Pause"
		request.LimitTime = strconv.Itoa(int(diff.Minutes() / 60))
		request.Percentage = strconv.Itoa(Percentage)
	} else {
		request.QuestEnd = time.Now().Format("2006-01-02 15:04:05")
		request.Status = "Failed"
		request.LimitTime = "0"
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
