package api

import (
	"fmt"
	cfg "game-play/internal/config"
	"game-play/src/infrastructure/restful/service/game_play"
	"game-play/src/infrastructure/router"
	container "game-play/src/shared/di"
	"game-play/src/shared/tracing"
	gameplayUsecase "game-play/src/usecase/game_play"
	Apis "github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"log"
)

func RunServer() {
	log.Println("Starting Restfull Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := Apis.New()
	ctn := container.NewContainer()

	tracer, closer := tracing.Init(e, "game-play", nil)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	Apply(e, ctn)
	svcPort := config.Server.Rest.Port

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *Apis.Echo, ctn *container.Container) {
	router.NewGamePlayRouter(e, game_play.NewGamePlayService(ctn.Resolve("game-play").(*gameplayUsecase.GamePlayInteractor)))
}
