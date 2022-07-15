package main

import (
	"fmt"
	"os"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/cmd/api"
	"github.com/p2034/HACK-2022-DREAM-DAIRY/cmd/auth"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Eror wrong argument count")
		return
	}

	whatWeDo := os.Args[1]
	if whatWeDo == "auth" {
		auth.Server()
	} else if whatWeDo == "api" {
		api.Server()
	} else {
		fmt.Println("Eror wrong argument: auth or api")
	}
}
