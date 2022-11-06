package model

type PlayersHistoryModel struct {
	BaseModels
	BaseCUModels
	PlayerId   string `gorm:"column:player_id"`
	LoginTime  string `gorm:"column:login_time"`
	LogoutTime string `gorm:"column:logout_time"`
}

func (PlayersHistoryModel) TableName() string {
	return "players_history"
}
