package model

type TasksModel struct {
	BaseModels
	BaseCUModels
	QuestId     string `gorm:"column:quest_id"`
	NameTask    string `gorm:"column:name_task"`
	Description string `gorm:"column:description"`
	ScoreTask   string `gorm:"column:score_task"`
}

func (TasksModel) TableName() string {
	return "task_model"
}
