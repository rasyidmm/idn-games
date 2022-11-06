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

type CreateQuestsRequest struct {
	NameQuest  string  `json:"name_quest"`
	LevelQuest string  `json:"level_quest"`
	TimeQuest  string  `json:"time_quest"`
	Tasks      []Tasks `json:"tasks"`
}
type CreateQuestsResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}
type GetQuestRequest struct {
	Id string `json:"id"  validate:"required"`
}

type Tasks struct {
	Id          string `json:"id"`
	NameTask    string `json:"name_task"`
	Description string `json:"description"`
	ScoreTask   string `json:"score_task"`
}

type Quest struct {
	Id         string  `json:"id"`
	NameQuest  string  `json:"name_quest"`
	LevelQuest string  `json:"level_quest"`
	TimeQuest  string  `json:"time_quest"`
	Tasks      []Tasks `json:"tasks"`
}

type ListQuestResponse struct {
	StatusCode string  `json:"status_code"`
	StatusDesc string  `json:"status_desc"`
	Quest      []Quest `json:"quest"`
}

type GetQuestResponse struct {
	Id        string  `json:"id"`
	NameQuest string  `json:"name_quest"`
	TimeQuest string  `json:"time_quest"`
	Tasks     []Tasks `json:"tasks"`
}
type QuestsClientHandler struct {
}

func NewQuestsClientHandler() *QuestsClientHandler {
	return &QuestsClientHandler{}
}

func (c *QuestsClientHandler) ListQuests(span opentracing.Span) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "ListQuests")
	defer sp.Finish()

	var (
		out *entity.ListQuestResponse
	)
	response := new(ListQuestResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.Games
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint

	restClient := client.SetRequest(endpoint, headers, nil)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", restClient.Body)

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
func (c *QuestsClientHandler) GetQuest(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetQuest")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetQuestRequest)
	var (
		out *entity.GetQuestResponse
	)
	response := new(GetQuestResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.Games
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint + "/" + reqdata.Id

	restClient := client.SetRequest(endpoint, headers, nil)

	tracing.LogObject(sp, "Endpoint", restClient.Endpoint)
	tracing.LogObject(sp, "Header", restClient.Headers)
	tracing.LogObject(sp, "Request", restClient.Body)
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
func (c *QuestsClientHandler) CreateQuest(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetPlayer")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.CreateQuestsRequest)
	var (
		request CreateQuestsRequest
		out     entity.CreateQuestsResponse
	)
	response := new(CreateQuestsResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, errors.New("request parsing err"))
		return nil, errors.New("request parsing err")
	}

	client := resty.New()
	uri := conf.Client.Games
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint

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
