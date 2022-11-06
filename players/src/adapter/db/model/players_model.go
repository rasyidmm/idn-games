package model

type PlayersModel struct {
	BaseModels
	BaseCUModels
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
	Nickname string `gorm:"column:nickname;unique"`
	Email    string `gorm:"column:email;unique"`
}

func (PlayersModel) TableName() string {
	return "players"
}
