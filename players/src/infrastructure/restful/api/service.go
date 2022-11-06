package api

import (
	"fmt"
	Apis "github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"log"
	cfg "players/internal/config"
	"players/src/infrastructure/restful/service/players"
	"players/src/infrastructure/router"
	playerUsecase "players/src/usecase/players"

	container "players/src/shared/di"

	"players/src/shared/tracing"
)

func RunServer() {
	log.Println("Starting Restfull Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := Apis.New()
	ctn := container.NewContainer()

	tracer, closer := tracing.Init(e, "players", nil)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	Apply(e, ctn)
	svcPort := config.Server.Rest.Port

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *Apis.Echo, ctn *container.Container) {
	router.NewPlayersRouter(e, players.NewPlayersService(ctn.Resolve("players").(*playerUsecase.PlayersInteractor)))
}
