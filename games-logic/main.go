package main

import (
	"fmt"
	restfulApi "games-logic/src/infrastructure/restful/api"
)

func main() {
	fmt.Println("Games - Logic")
	restfulApi.RunServer()
}
