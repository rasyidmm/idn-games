package dto

import "github.com/mitchellh/mapstructure"

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
type GamePlayBuilder struct {
}

func (b *GamePlayBuilder) StartGameResponse(in interface{}) (interface{}, error) {
	var out *StartGameResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *GamePlayBuilder) EndGameResponse(in interface{}) (interface{}, error) {
	var out *EndGameResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *GamePlayBuilder) ListGamePlayByPlayerIdResponse(in interface{}) (interface{}, error) {
	var out *ListGamePlayByPlayerIdResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *GamePlayBuilder) GetGamePlayResponse(in interface{}) (interface{}, error) {
	var out *GetGamePlayResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *GamePlayBuilder) PauseGameResponse(in interface{}) (interface{}, error) {
	var out *PauseGameResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
