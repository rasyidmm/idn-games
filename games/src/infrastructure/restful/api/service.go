package api

import (
	"fmt"
	cfg "games/internal/config"
	"games/src/infrastructure/restful/service/quests_service"
	"games/src/infrastructure/router"
	container "games/src/shared/di"
	"games/src/shared/tracing"
	"games/src/usecase/quests"
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

	tracer, closer := tracing.Init(e, "games", nil)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	Apply(e, ctn)
	svcPort := config.Server.Rest.Port

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *Apis.Echo, ctn *container.Container) {
	router.NewQuestsRouter(e, quests_service.NewQuestsService(ctn.Resolve("quests").(*quests.QuestInteractor)))
}
