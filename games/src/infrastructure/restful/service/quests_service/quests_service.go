package quests_service

import (
	"games/src/shared/tracing"
	usecase "games/src/usecase/quests"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type CreateQuestsRequest struct {
	NameQuest  string  `json:"name_quest"  validate:"required"`
	LevelQuest string  `json:"level_quest"  validate:"required"`
	TimeQuest  string  `json:"time_quest"  validate:"required"`
	Tasks      []Tasks `json:"tasks"  validate:"required"`
}
type Tasks struct {
	NameTask    string `json:"name_task"  validate:"required"`
	Description string `json:"description"  validate:"required"`
	ScoreTask   string `json:"score_task"  validate:"required"`
}

type GetQuestRequest struct {
	Id string `json:"id"  validate:"required"`
}

type QuestsService struct {
	uc usecase.QuestsInputPort
}

func NewQuestsService(u usecase.QuestsInputPort) *QuestsService {
	return &QuestsService{
		uc: u,
	}
}

func (s *QuestsService) CreateQuest(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "CreateQuest")
	defer sp.Finish()

	reqData := new(CreateQuestsRequest)
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

	var request *usecase.CreateQuestsRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.CreateQuest(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}

func (s *QuestsService) GetQuest(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "GetQuest")
	defer sp.Finish()

	id := ctx.Param("id")
	reqData := &GetQuestRequest{id}
	tracing.LogRequest(sp, reqData)

	err := ctx.Validate(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var request *usecase.GetQuestRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := s.uc.GetQuest(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}

func (s *QuestsService) ListQuests(ctx echo.Context) error {
	sp, _ := tracing.CreateRootSpan(ctx, "ListQuests")
	defer sp.Finish()

	res, err := s.uc.ListQuests(ctx)
	if err != nil {
		tracing.LogError(sp, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogResponse(sp, res)
	return ctx.JSON(http.StatusOK, res)
}
