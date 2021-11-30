package cmd

import (
	"fmt"
	"projectstructuring/configs"
	"projectstructuring/routes"
)

func Run() {
	fmt.Println("Entering cmd")
	configs.InitDB(configs.SetENV())
	routes.Route()
}
