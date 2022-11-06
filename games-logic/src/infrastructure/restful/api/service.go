package api

import (
	"fmt"
	cfg "games-logic/internal/config"
	"games-logic/src/infrastructure/restful/service/game_play"
	"games-logic/src/infrastructure/restful/service/players"
	"games-logic/src/infrastructure/restful/service/quests_service"
	"games-logic/src/infrastructure/router"
	container "games-logic/src/shared/di"
	gamePlayUsecase "games-logic/src/usecase/game_play"
	playerUsecase "games-logic/src/usecase/players"
	"games-logic/src/usecase/quests"
	Apis "github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"log"

	"games-logic/src/shared/tracing"
)

func RunServer() {
	log.Println("Starting Restfull Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := Apis.New()
	ctn := container.NewContainer()

	tracer, closer := tracing.Init(e, "games-logic", nil)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	Apply(e, ctn)
	svcPort := config.Server.Rest.Port

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *Apis.Echo, ctn *container.Container) {
	router.NewPlayersRouter(e, players.NewPlayersService(ctn.Resolve("players").(*playerUsecase.PlayersInteractor)))
	router.NewGamePlayRouter(e, game_play.NewGamePlayService(ctn.Resolve("game-play").(*gamePlayUsecase.GamePlayInteractor)))
	router.NewQuestsRouter(e, quests_service.NewQuestsService(ctn.Resolve("quest").(*quests.QuestInteractor)))
}
