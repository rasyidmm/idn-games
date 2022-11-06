package client

import (
	"errors"
	cfg "games-logic/internal/config"
	"games-logic/src/adapter/resty"
	"games-logic/src/domain/entity"
	"games-logic/src/shared/tracing"
	"github.com/mitchellh/mapstructure"
	"github.com/opentracing/opentracing-go"
)

type StartGameRequest struct {
	Id        string `json:"id"`
	PlayerId  string `json:"player_id"`
	QuestId   string `json:"quest_id"`
	Status    string `json:"status"`
	LimitTime string `json:"limit_time"`
}
type EndGameRequest struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
	QuestEnd   string `json:"quest_end"`
	Status     string `json:"status"`
	Score      string `json:"score"`
	LimitTime  string `json:"limit_time"`
	Percentage string `json:"percentage"`
}
type ListGamePlayByPlayerIdRequest struct {
	PlayerId string `json:"player_id"`
}

type StartGameResponse struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	QuestStart string `json:"quest_start"`
}

type EndGameResponse struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
	QuestStart string `json:"quest_start"`
	QuestEnd   string `json:"quest_end"`
	Status     string `json:"status"`
	Score      string `json:"score"`
	LimitTime  string `json:"limit_time"`
	Percentage string `json:"percentage"`
}
type GamePlay struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
	QuestStart string `json:"quest_start"`
	QuestEnd   string `json:"quest_end"`
	Status     string `json:"status"`
	Score      string `json:"score"`
	LimitTime  string `json:"limit_time"`
	Percentage string `json:"percentage"`
}

type ListGamePlayByPlayerIdResponse struct {
	GamePlay []GamePlay `json:"game_play"`
}

type GetGamePlayRequest struct {
	Id string `json:"id"  validate:"required"`
}

type GetGamePlayResponse struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
	QuestStart string `json:"quest_start"`
	QuestEnd   string `json:"quest_end"`
	Status     string `json:"status"`
	Score      string `json:"score"`
	LimitTime  string `json:"limit_time"`
}

type PauseGameRequest struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
	QuestEnd   string `json:"quest_end"`
	Status     string `json:"status"`
	Score      string `json:"score"`
	LimitTime  string `json:"limit_time"`
	Percentage string `json:"percentage"`
}
type PauseGameResponse struct {
	Id         string `json:"id"`
	PlayerId   string `json:"player_id"`
	QuestId    string `json:"quest_id"`
	ListTaskId string `json:"list_task_id"`
	QuestStart string `json:"quest_start"`
	QuestEnd   string `json:"quest_end"`
	Status     string `json:"status"`
	Score      string `json:"score"`
	LimitTime  string `json:"limit_time"`
	Percentage string `json:"percentage"`
}
type GamePlayClientHandler struct {
}

func NewGamePlayClientHandler() *GamePlayClientHandler {
	return &GamePlayClientHandler{}
}

func (c GamePlayClientHandler) ListGamePlayByPlayerId(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "ListGamePlayByPlayerId")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.ListGamePlayByPlayerIdRequest)
	var (
		request ListGamePlayByPlayerIdRequest
		out     *entity.ListGamePlayByPlayerIdResponse
	)
	response := new(ListGamePlayByPlayerIdResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.GamePlay
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint + "/list-by-player-id"

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	restClient := client.SetRequest(endpoint, headers, request)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", restClient.Body)

	err = restClient.Post(&response)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	err = mapstructure.Decode(response, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return out, nil
}

func (c GamePlayClientHandler) StartGame(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "StartGame")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.StartGameRequest)
	var (
		request StartGameRequest
		out     *entity.StartGameResponse
	)
	response := new(StartGameResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.GamePlay
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint + "/start-game"

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	restClient := client.SetRequest(endpoint, headers, request)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", restClient.Body)

	err = restClient.Post(&response)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	err = mapstructure.Decode(response, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return out, nil
}

func (c GamePlayClientHandler) EndGame(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "EndGame")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.EndGameRequest)
	var (
		request EndGameRequest
		out     *entity.EndGameResponse
	)
	response := new(EndGameResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.GamePlay
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint + "/end-game"

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	restClient := client.SetRequest(endpoint, headers, request)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", restClient.Body)

	err = restClient.Post(&response)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	err = mapstructure.Decode(response, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return out, nil
}

func (c *GamePlayClientHandler) GetGamePlay(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetGamePlay")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetGamePlayRequest)
	var (
		out *entity.GetGamePlayResponse
	)
	response := new(GetGamePlayResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.GamePlay
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint + "/" + reqdata.Id

	restClient := client.SetRequest(endpoint, headers, nil)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", reqdata.Id)
	err := restClient.Get(&response)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	err = mapstructure.Decode(response, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return out, nil
}
func (c GamePlayClientHandler) PauseGame(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PauseGame")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.PauseGameRequest)
	var (
		request PauseGameRequest
		out     *entity.PauseGameResponse
	)
	response := new(PauseGameResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.GamePlay
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint + "/pause-game"

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	restClient := client.SetRequest(endpoint, headers, request)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", restClient.Body)

	err = restClient.Post(&response)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	err = mapstructure.Decode(response, &out)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	tracing.LogResponse(sp, out)
	return out, nil
}
