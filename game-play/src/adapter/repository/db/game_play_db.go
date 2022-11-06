package db

import (
	"game-play/src/adapter/db/model"
	"game-play/src/domain/entity"
	"game-play/src/shared/tracing"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type GamePlayDataHandle struct {
	db *gorm.DB
}

func NewGamePlayDataHandle(db *gorm.DB) *GamePlayDataHandle {
	return &GamePlayDataHandle{
		db: db,
	}
}

func (d *GamePlayDataHandle) StartGame(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "StartGame")
	defer sp.Finish()
	tracing.LogRequest(sp, in)
	var datas model.GamePlayModel

	reqdata := in.(*entity.StartGameRequest)

	if reqdata.Id != "" {
		err := d.db.Debug().Where("id = ?", reqdata.Id).First(&datas).Error
		if err != nil {
			tracing.LogError(sp, err)
			return nil, err
		}
	}

	datas.CreatedAt = time.Now()
	datas.UpdatedAt = time.Now()
	datas.PlayerId = reqdata.PlayerId
	datas.QuestId = reqdata.QuestId
	datas.QuestStart = time.Now().Format("2006-01-02 15:04:05")
	datas.StatusGame = reqdata.Status
	if reqdata.Id == "" {
		datas.LimitTime = reqdata.LimitTime
	}

	tracing.LogObject(sp, "data entity", datas)
	err := d.db.Debug().Save(&datas).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.StartGameResponse{
		Id:         strconv.Itoa(int(datas.Id)),
		PlayerId:   datas.PlayerId,
		QuestId:    datas.QuestId,
		QuestStart: datas.QuestStart,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
func (d *GamePlayDataHandle) EndGame(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "EndGame")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.EndGameRequest)
	var data model.GamePlayModel

	tracing.LogObject(sp, "data entity", data)
	err := d.db.Debug().Where("id = ?", reqdata.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	data.QuestEnd = reqdata.QuestEnd
	data.ListTaskId = reqdata.ListTaskId
	data.ScoreQuest = reqdata.Score
	data.StatusGame = reqdata.Status
	data.LimitTime = reqdata.LimitTime
	data.Percentage = reqdata.Percentage

	err = d.db.Debug().Where("id = ?", reqdata.Id).Save(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.EndGameResponse{
		Id:         strconv.Itoa(int(data.Id)),
		PlayerId:   data.PlayerId,
		QuestId:    data.QuestId,
		ListTaskId: data.ListTaskId,
		QuestStart: data.QuestStart,
		QuestEnd:   data.QuestEnd,
		Score:      data.ScoreQuest,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *GamePlayDataHandle) ListGamePlayByPlayerId(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "ListGamePlayByPlayerId")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.ListGamePlayByPlayerIdRequest)
	var (
		data    []model.GamePlayModel
		resData []entity.GamePlay
	)
	err := d.db.Debug().Where("player_id = ?", reqdata.PlayerId).Find(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	for _, item := range data {
		resData = append(resData, entity.GamePlay{
			Id:         strconv.Itoa(int(item.Id)),
			PlayerId:   item.PlayerId,
			QuestId:    item.QuestId,
			ListTaskId: item.ListTaskId,
			QuestStart: item.QuestStart,
			QuestEnd:   item.QuestEnd,
			Status:     item.StatusGame,
			Score:      item.ScoreQuest,
			LimitTime:  item.LimitTime,
			Percentage: item.Percentage,
		})
	}

	res := &entity.ListGamePlayByPlayerIdResponse{GamePlay: resData}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *GamePlayDataHandle) GetGamePlay(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetGamePlay")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetGamePlayRequest)
	var (
		data model.GamePlayModel
	)
	err := d.db.Debug().Where("id = ?", reqdata.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res := &entity.GetGamePlayResponse{
		Id:         strconv.Itoa(int(data.Id)),
		PlayerId:   data.PlayerId,
		QuestId:    data.QuestId,
		ListTaskId: data.ListTaskId,
		QuestStart: data.QuestStart,
		QuestEnd:   data.QuestId,
		Status:     data.StatusGame,
		Score:      data.ScoreQuest,
		LimitTime:  data.LimitTime,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *GamePlayDataHandle) PauseGame(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PauseGame")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.PauseGameRequest)
	var data model.GamePlayModel

	tracing.LogObject(sp, "data entity", data)
	err := d.db.Debug().Where("id = ?", reqdata.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	data.ListTaskId = reqdata.ListTaskId
	data.ScoreQuest = reqdata.Score
	data.StatusGame = reqdata.Status
	data.LimitTime = reqdata.LimitTime
	data.Percentage = reqdata.Percentage
	data.QuestStart = ""

	err = d.db.Debug().Where("id = ?", reqdata.Id).Save(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.PauseGameResponse{
		Id:         strconv.Itoa(int(data.Id)),
		PlayerId:   data.PlayerId,
		QuestId:    data.QuestId,
		ListTaskId: data.ListTaskId,
		QuestStart: data.QuestStart,
		QuestEnd:   data.QuestEnd,
		Score:      data.ScoreQuest,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
