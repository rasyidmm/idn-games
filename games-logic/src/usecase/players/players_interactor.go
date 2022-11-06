package players

import (
	"errors"
	"games-logic/src/domain/entity"
	"games-logic/src/domain/repository"
	"games-logic/src/shared/enum"
	"games-logic/src/shared/tracing"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"strings"
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
	Gold     string
	GamePlay []GamePlayPlayer
}
type GamePlayPlayer struct {
	Id         string
	QuestStart string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
	Quest      []Quest
}
type Tasks struct {
	Id          string
	NameTask    string
	Description string
	ScoreTask   string
}

type Quest struct {
	Id         string
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}

type PlayersInteractor struct {
	repo   repository.PlayersRepository
	repoG  repository.QuestsRepository
	repoGp repository.GamePlayRepository
	out    PlayerOutputPort
}

func NewPlayersInteractor(r repository.PlayersRepository, rg repository.QuestsRepository, rgp repository.GamePlayRepository, o PlayerOutputPort) *PlayersInteractor {
	return &PlayersInteractor{
		repo:   r,
		repoG:  rg,
		repoGp: rgp,
		out:    o,
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
	outPlayer := resData.(*entity.GetPlayerResponse)

	reqGp := &entity.ListGamePlayByPlayerIdRequest{PlayerId: outPlayer.Id}
	resGp, err := u.repoGp.ListGamePlayByPlayerId(sp, reqGp)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	outGamePlay := resGp.(*entity.ListGamePlayByPlayerIdResponse)

	resG, err := u.repoG.ListQuests(sp)

	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	outGame := resG.(*entity.ListQuestResponse)

	var (
		GamePlay []GamePlayPlayer
		gold     int
	)
	for _, itemGP := range outGamePlay.GamePlay {
		if itemGP.Status != "Failed" {
			var quests []Quest
			if itemGP.Status == "Done" {
				gordItem, _ := strconv.Atoi(itemGP.Score)
				gold += gordItem
			}
			for _, itemG := range outGame.Quest {
				var tasks []Tasks
				if itemGP.QuestId == itemG.Id {
					listTaksId := strings.Split(itemGP.ListTaskId, ";")
					for _, itemTask := range listTaksId {
						for _, task := range itemG.Tasks {
							if itemTask == task.Id {
								tasks = append(tasks, Tasks{
									Id:          task.Id,
									NameTask:    task.NameTask,
									Description: task.Description,
									ScoreTask:   task.ScoreTask,
								})
							}
						}
					}
					quests = append(quests, Quest{
						Id:         itemG.Id,
						NameQuest:  itemG.NameQuest,
						LevelQuest: itemG.LevelQuest,
						TimeQuest:  itemG.TimeQuest,
						Tasks:      tasks,
					})
				}
			}
			GamePlay = append(GamePlay, GamePlayPlayer{
				Id:         itemGP.Id,
				QuestStart: itemGP.QuestStart,
				QuestEnd:   itemGP.QuestEnd,
				Status:     itemGP.Status,
				Score:      itemGP.Score,
				LimitTime:  itemGP.LimitTime,
				Percentage: itemGP.Percentage,
				Quest:      quests,
			})
		}

	}

	var res *GetPlayerResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res.GamePlay = GamePlay
	res.Gold = strconv.Itoa(gold)

	tracing.LogResponse(sp, res)
	return u.out.GetPlayerResponse(res)
}
