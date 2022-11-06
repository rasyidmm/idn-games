package di

import (
	"games-logic/src/adapter/client"
	"games-logic/src/infrastructure/restful/service/dto"
	"games-logic/src/usecase/game_play"
	"games-logic/src/usecase/players"
	"games-logic/src/usecase/quests"
	"github.com/sarulabs/di"
)

// Container :
type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "players", Build: playerUsecase},
		{Name: "game-play", Build: gamePlayUsecase},
		{Name: "quest", Build: questsUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func playerUsecase(_ di.Container) (interface{}, error) {
	repo := client.NewPlayerClientHandler()
	repog := client.NewQuestsClientHandler()
	repogp := client.NewGamePlayClientHandler()
	out := &dto.PlayersBuilder{}
	return players.NewPlayersInteractor(repo, repog, repogp, out), nil
}

func gamePlayUsecase(_ di.Container) (interface{}, error) {
	repo := client.NewGamePlayClientHandler()
	repog := client.NewQuestsClientHandler()
	out := &dto.GamePlayBuilder{}
	return game_play.NewGamePlayInteractor(repo, repog, out), nil
}

func questsUsecase(_ di.Container) (interface{}, error) {
	repo := client.NewQuestsClientHandler()
	out := &dto.QuestBuilder{}
	return quests.NewQuestInteractor(repo, out), nil
}
