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

type CreatePlayerRequest struct {
	Username string `json:"username"`
	Password string `json:"password" `
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type CreatePlayerResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type GetPlayerResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type PlayerClientHandler struct {
}

func NewPlayerClientHandler() *PlayerClientHandler {
	return &PlayerClientHandler{}
}

func (c *PlayerClientHandler) CreatePlayer(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreatePlayer")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.CreatePlayerRequest)
	var (
		request CreatePlayerRequest
		out     entity.CreatePlayerResponse
	)
	response := new(CreatePlayerResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.Players
	endpoint := uri.Url + ":" + uri.Port + "/" + uri.Endpoint

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

func (c *PlayerClientHandler) GetPlayer(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetPlayer")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetPlayerRequest)
	var (
		out *entity.GetPlayerResponse
	)
	response := new(GetPlayerResponse)

	conf := cfg.GetConfig()
	headers := map[string]string{
		"Content-type": "application/json",
	}

	client := resty.New()
	uri := conf.Client.Players
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
