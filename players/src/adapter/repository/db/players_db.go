package db

import (
	"github.com/mitchellh/mapstructure"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"players/src/adapter/db/model"
	"players/src/domain/entity"
	"players/src/shared/tracing"
	"players/src/shared/util"
	"strconv"
	"time"
)

type PlayersDataHandler struct {
	db *gorm.DB
}

func NewPlayersDataHandler(db *gorm.DB) *PlayersDataHandler {
	return &PlayersDataHandler{
		db: db,
	}
}

func (d *PlayersDataHandler) CreatePlayer(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreatePlayer")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.CreatePlayerRequest)
	data := &model.PlayersModel{
		BaseModels: model.BaseModels{
			CreateBy: reqdata.Username,
		},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()},
		Username: reqdata.Username,
		Password: util.HashSha512(reqdata.Password),
		Nickname: reqdata.Nickname,
		Email:    reqdata.Email,
	}

	tracing.LogObject(sp, "data entity", data)
	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.CreatePlayerResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *PlayersDataHandler) GetPlayer(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetPlayer")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetPlayerRequest)
	var (
		data model.PlayersModel
		res  entity.GetPlayerResponse
	)

	err := d.db.Debug().Where("id = ?", reqdata.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	err = mapstructure.Decode(data, &res)
	res.Id = strconv.Itoa(int(data.Id))
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
