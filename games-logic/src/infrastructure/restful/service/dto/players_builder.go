package dto

import "github.com/mitchellh/mapstructure"

type CreatePlayerResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type GetPlayerResponse struct {
	Id       string           `json:"id"`
	Username string           `json:"username"`
	Nickname string           `json:"nickname"`
	Email    string           `json:"email"`
	Gold     string           `json:"gold"`
	GamePlay []GamePlayPlayer `json:"game_play"`
}
type GamePlayPlayer struct {
	Id         string  `json:"id"`
	QuestStart string  `json:"quest_start"`
	QuestEnd   string  `json:"quest_end"`
	Status     string  `json:"status"`
	Score      string  `json:"score"`
	LimitTime  string  `json:"limit_time"`
	Percentage string  `json:"percentage"`
	Quest      []Quest `json:"quest"`
}
type PlayersBuilder struct {
}

func (b *PlayersBuilder) CreatePlayerResponse(in interface{}) (interface{}, error) {
	var out *CreatePlayerResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *PlayersBuilder) GetPlayerResponse(in interface{}) (interface{}, error) {
	var out *GetPlayerResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
