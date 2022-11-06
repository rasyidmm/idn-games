package entity

type CreateQuestsRequest struct {
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}
type CreateQuestsResponse struct {
	StatusCode string
	StatusDesc string
}
type CreateTasksRequest struct {
	QuestId     string
	NameTask    string
	Description string
	ScoreTask   string
}

type CreateTasksResponse struct {
	StatusCode string
	StatusDesc string
}

type Quest struct {
	Id         string
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}

type ListQuestResponse struct {
	StatusCode string
	StatusDesc string
	Quest      []Quest
}
type GetQuestRequest struct {
	Id string
}
type GetQuestResponse struct {
	Id         string
	NameQuest  string
	LevelQuest string
	TimeQuest  string
	Tasks      []Tasks
}
type GetTaskRequest struct {
	Id string
}

type GetTaskResponse struct {
	Id          string
	QuestId     string
	NameTask    string
	Description string
	ScoreTask   string
}
type Tasks struct {
	Id          string
	QuestId     string
	NameTask    string
	Description string
	ScoreTask   string
}

type ListTasksResponse struct {
	StatusCode string
	StatusDesc string
	Tasks      []Tasks
}

type ListTaskByQuestId struct {
	Id string
}
