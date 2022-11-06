package quests

import (
	"errors"
	"games/src/domain/entity"
	"games/src/domain/repository"
	"games/src/shared/enum"
	"games/src/shared/tracing"
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
	repo repository.QuestRepository
	out  QuestsOutputPort
}

func NewQuestInteractor(r repository.QuestRepository, o QuestsOutputPort) *QuestInteractor {
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

	var (
		quests []Quest
	)

	resListQuest, err := u.repo.ListQuest(sp)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	outListQuest := resListQuest.(*entity.ListQuestResponse)

	resListTasks, err := u.repo.ListTask(sp)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	outTaskQuest := resListTasks.(*entity.ListTasksResponse)

	for _, itemQuest := range outListQuest.Quest {
		var tasks []Tasks
		for _, itemTask := range outTaskQuest.Tasks {
			if itemTask.QuestId == itemQuest.Id {
				tasks = append(tasks, Tasks{
					Id:          itemTask.Id,
					QuestId:     itemTask.QuestId,
					NameTask:    itemTask.NameTask,
					Description: itemTask.Description,
					ScoreTask:   itemTask.ScoreTask,
				})
			}
		}

		quests = append(quests, Quest{
			Id:         itemQuest.Id,
			NameQuest:  itemQuest.NameQuest,
			LevelQuest: itemQuest.LevelQuest,
			TimeQuest:  itemQuest.TimeQuest,
			Tasks:      tasks,
		})
	}

	out := ListQuestResponse{
		StatusCode: outListQuest.StatusCode,
		StatusDesc: outListQuest.StatusDesc,
		Quest:      quests,
	}

	tracing.LogResponse(sp, out)
	return u.out.ListQuestResponse(out)
}
