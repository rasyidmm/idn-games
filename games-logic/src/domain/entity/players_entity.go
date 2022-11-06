package entity

type CreatePlayerRequest struct {
	Username string
	Password string
	Nickname string
	Email    string
}

type CreatePlayerResponse struct {
	StatusCode string
	StatusDesc string
}

type GetPlayerRequest struct {
	Id string
}
type GetPlayerResponse struct {
	Id       string
	Username string
	Password string
	Nickname string
	Email    string
}

type GetGamePlayRequest struct {
	Id string
}

type GetGamePlayResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
}
