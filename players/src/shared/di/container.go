package di

import (
	"github.com/sarulabs/di"
	"players/src/adapter/db/connection"
	"players/src/adapter/repository/db"
	"players/src/infrastructure/restful/service/dto"
	"players/src/usecase/players"
)

// Container :
type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "players", Build: playerUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func playerUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewPlayersDataHandler(connection.GamesDB)
	out := &dto.PlayersBuilder{}
	return players.NewPlayersInteractor(repo, out), nil
}
