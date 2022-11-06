package di

import (
	"games/src/adapter/db/connection"
	"games/src/adapter/repository/db"
	"games/src/infrastructure/restful/service/dto"
	"games/src/usecase/quests"
	"github.com/sarulabs/di"
)

// Container :
type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "quests", Build: questUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func questUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewQuestsDataHandler(connection.GamesDB)
	out := &dto.QuestBuilder{}
	return quests.NewQuestInteractor(repo, out), nil
}
