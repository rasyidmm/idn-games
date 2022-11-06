package di

import (
	"game-play/src/adapter/db/connection"
	"game-play/src/adapter/repository/db"
	"game-play/src/infrastructure/restful/service/dto"
	"game-play/src/usecase/game_play"
	"github.com/sarulabs/di"
)

// Container :
type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "game-play", Build: gamePlayUscase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}
func gamePlayUscase(_ di.Container) (interface{}, error) {
	repo := db.NewGamePlayDataHandle(connection.GamesDB)
	out := &dto.GamePlayBuilder{}
	return game_play.NewGamePlayInteractor(repo, out), nil
}
