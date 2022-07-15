package main

import (
	"fmt"
	"os"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/cmd/api"
	"github.com/p2034/HACK-2022-DREAM-DAIRY/cmd/auth"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Eror wrong argument count")
		return
	}

	whatWeDo := os.Args[1]
	if whatWeDo == "auth" {
		auth.Server(os.Args[2])
	} else if whatWeDo == "api" {
		api.Server(os.Args[2])
	} else {
		fmt.Println("Error wrong argument: auth or api")
	}
}
