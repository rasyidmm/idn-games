package main

import (
	"fmt"
	restfulApi "game-play/src/infrastructure/restful/api"
)

func main() {
	fmt.Println("Game-Play")
	restfulApi.RunServer()
}
