package entity

type StartGameRequest struct {
	Id        string
	PlayerId  string
	QuestId   string
	Status    string
	LimitTime string
}
type StartGameResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	QuestStart string
}
type EndGameRequest struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type EndGameResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Score      string
	LimitTime  string
	Percentage string
}
type ListGamePlayByPlayerIdRequest struct {
	PlayerId string
}

type GamePlay struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type ListGamePlayByPlayerIdResponse struct {
	GamePlay []GamePlay
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
	Percentage string
}

type PauseGameRequest struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestEnd   string
	Status     string
	Score      string
	LimitTime  string
	Percentage string
}

type PauseGameResponse struct {
	Id         string
	PlayerId   string
	QuestId    string
	ListTaskId string
	QuestStart string
	QuestEnd   string
	Score      string
	LimitTime  string
	Percentage string
}
