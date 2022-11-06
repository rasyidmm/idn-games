package db

import (
	"games/src/adapter/db/model"
	"games/src/domain/entity"
	"games/src/shared/tracing"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type QuestsDataHandler struct {
	db *gorm.DB
}

func NewQuestsDataHandler(db *gorm.DB) *QuestsDataHandler {
	return &QuestsDataHandler{
		db: db,
	}
}

func (d *QuestsDataHandler) CreateQuest(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateQuest")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.CreateQuestsRequest)
	var tasks []model.TasksModel

	tx := d.db.Debug()
	data := &model.QuestsModel{
		BaseModels: model.BaseModels{},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LevelQuest: reqdata.LevelQuest,
		NameQuest:  reqdata.NameQuest,
		TimeQuest:  reqdata.TimeQuest,
	}
	tracing.LogObject(sp, "Quests entity", data)
	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		tracing.LogError(sp, err)
		return nil, err
	}

	for _, item := range reqdata.Tasks {
		tasks = append(tasks, model.TasksModel{
			BaseCUModels: model.BaseCUModels{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			QuestId:      strconv.Itoa(int(data.Id)),
			NameTask:     item.NameTask,
			Description:  item.Description,
			ScoreTask:    item.ScoreTask,
		})
	}

	tracing.LogObject(sp, "tasks entity", tasks)
	err = tx.Create(&tasks).Error
	if err != nil {
		tx.Rollback()
		tracing.LogError(sp, err)
		return nil, err
	}

	tx.Commit()

	res := &entity.CreateQuestsResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil

}

func (d *QuestsDataHandler) CreateTask(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateTask")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.CreateTasksRequest)

	data := &model.TasksModel{
		BaseModels: model.BaseModels{},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		QuestId:     reqdata.QuestId,
		NameTask:    reqdata.NameTask,
		Description: reqdata.Description,
		ScoreTask:   reqdata.ScoreTask,
	}

	tracing.LogObject(sp, "data entity", data)
	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.CreateQuestsResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *QuestsDataHandler) GetQuest(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetQuest")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetQuestRequest)

	var (
		data        model.QuestsModel
		tasks       []model.TasksModel
		tasksEntity []entity.Tasks
	)

	err := d.db.Debug().Where("id = ?", reqdata.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	err = d.db.Debug().Where("quest_id = ?", reqdata.Id).Find(&tasks).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	for _, item := range tasks {
		tasksEntity = append(tasksEntity, entity.Tasks{
			Id:          strconv.Itoa(int(item.Id)),
			QuestId:     item.QuestId,
			NameTask:    item.NameTask,
			Description: item.Description,
			ScoreTask:   item.ScoreTask,
		})
	}

	res := &entity.GetQuestResponse{
		Id:         strconv.Itoa(int(data.Id)),
		LevelQuest: data.LevelQuest,
		NameQuest:  data.NameQuest,
		TimeQuest:  data.TimeQuest,
		Tasks:      tasksEntity,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *QuestsDataHandler) ListQuest(span opentracing.Span) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetQuest")
	defer sp.Finish()

	var (
		data  []model.QuestsModel
		quest []entity.Quest
	)

	err := d.db.Debug().Find(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	for _, item := range data {
		quest = append(quest, entity.Quest{
			Id:         strconv.Itoa(int(item.Id)),
			NameQuest:  item.NameQuest,
			LevelQuest: item.LevelQuest,
			TimeQuest:  item.TimeQuest,
		})
	}

	res := &entity.ListQuestResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
		Quest:      quest,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *QuestsDataHandler) GetTask(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetTask")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetTaskRequest)

	var (
		data model.TasksModel
	)

	err := d.db.Debug().Where("id = ?", reqdata.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.GetTaskResponse{
		Id:          strconv.Itoa(int(data.Id)),
		QuestId:     data.QuestId,
		NameTask:    data.NameTask,
		Description: data.Description,
		ScoreTask:   data.ScoreTask,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *QuestsDataHandler) ListTask(span opentracing.Span) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "ListTask")
	defer sp.Finish()

	var (
		data  []model.TasksModel
		tasks []entity.Tasks
	)

	err := d.db.Debug().Find(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	for _, item := range data {
		tasks = append(tasks, entity.Tasks{
			Id:          strconv.Itoa(int(item.Id)),
			QuestId:     item.QuestId,
			NameTask:    item.NameTask,
			Description: item.Description,
			ScoreTask:   item.ScoreTask,
		})
	}

	res := &entity.ListTasksResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
		Tasks:      tasks,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *QuestsDataHandler) ListTaskByQuestId(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "ListTaskByQuestId")
	defer sp.Finish()
	reqData := in.(*entity.ListTaskByQuestId)

	var (
		data  []model.TasksModel
		tasks []entity.Tasks
	)

	err := d.db.Debug().Where("quest_id = ?", reqData.Id).Find(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	for _, item := range data {
		tasks = append(tasks, entity.Tasks{
			Id:          strconv.Itoa(int(item.Id)),
			QuestId:     item.QuestId,
			NameTask:    item.NameTask,
			Description: item.Description,
			ScoreTask:   item.ScoreTask,
		})
	}

	res := &entity.ListTasksResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
		Tasks:      tasks,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
