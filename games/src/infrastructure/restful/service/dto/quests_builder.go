package dto

import "github.com/mitchellh/mapstructure"

type CreateQuestsResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
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

type QuestBuilder struct {
}

func (b *QuestBuilder) CreateQuestsResponse(in interface{}) (interface{}, error) {
	var out *CreateQuestsResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *QuestBuilder) GetQuestResponse(in interface{}) (interface{}, error) {
	var out *GetQuestResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *QuestBuilder) ListQuestResponse(in interface{}) (interface{}, error) {
	var out *ListQuestResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
