package quests

import (
	"errors"
	"games-logic/src/domain/entity"
	"games-logic/src/domain/repository"
	"games-logic/src/shared/enum"
	"games-logic/src/shared/tracing"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type CreateQuestsRequest struct {
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}
type Tasks struct {
	Id          string
	QuestId     string
	NameTask    string
	Description string
	ScoreTask   string
}

type CreateQuestsResponse struct {
	StatusCode string
	StatusDesc string
}

type Quest struct {
	Id         string
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}

type ListQuestResponse struct {
	StatusCode string
	StatusDesc string
	Quest      []Quest
}
type GetQuestRequest struct {
	Id string
}
type GetQuestResponse struct {
	Id         string
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}

type QuestInteractor struct {
	repo repository.QuestsRepository
	out  QuestsOutputPort
}

func NewQuestInteractor(r repository.QuestsRepository, o QuestsOutputPort) *QuestInteractor {
	return &QuestInteractor{
		repo: r,
		out:  o,
	}
}

func (u *QuestInteractor) CreateQuest(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*CreateQuestsRequest)
	var request *entity.CreateQuestsRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	res, err := u.repo.CreateQuest(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var out CreateQuestsResponse
	err = mapstructure.Decode(res, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return u.out.CreateQuestsResponse(out)
}

func (u *QuestInteractor) GetQuest(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)
	reqdata := in.(*GetQuestRequest)

	var request *entity.GetQuestRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}
	res, err := u.repo.GetQuest(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	var out GetQuestResponse
	err = mapstructure.Decode(res, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return u.out.GetQuestResponse(out)

}
func (u *QuestInteractor) ListQuests(ctx echo.Context) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()

	res, err := u.repo.ListQuests(sp)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var out ListQuestResponse
	err = mapstructure.Decode(res, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return u.out.ListQuestResponse(out)
}
