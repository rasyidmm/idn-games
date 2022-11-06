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
