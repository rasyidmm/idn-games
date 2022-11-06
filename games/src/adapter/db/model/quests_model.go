package model

type QuestsModel struct {
	BaseModels
	BaseCUModels
	NameQuest  string `gorm:"column:name_quest"`
	LevelQuest string `gorm:"column:level_quest"`
	TimeQuest  string `gorm:"column:time_quest"`
}

func (QuestsModel) TableName() string {
	return "quests_model"
}
