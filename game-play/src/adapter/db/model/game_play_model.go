package model

type GamePlayModel struct {
	BaseModels
	BaseCUModels
	PlayerId   string `gorm:"column:player_id"`
	QuestId    string `gorm:"column:quest_id"`
	ListTaskId string `gorm:"column:list_task_id"`
	QuestStart string `gorm:"column:quest_start"`
	QuestEnd   string `gorm:"column:quest_end"`
	LimitTime  string `gorm:"column:limit_time"`
	StatusGame string `gorm:"column:status_game"`
	ScoreQuest string `gorm:"column:score_quest"`
	Percentage string `gorm:"column:percentage"`
}

func (GamePlayModel) TableName() string {
	return "game_play_model"
}
