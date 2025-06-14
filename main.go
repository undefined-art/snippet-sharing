package main

import (
	"flag"
	"fmt"
	"gin-rest/config"
	"gin-rest/internal/routes"

	"os"
)

func main() {
	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	// TODO: Implement DB initialization
	// db.Init()

	Init()
}

func Init() {
	config := config.GetConfig()
	r := routes.NewRouter()
	config.GetString("server.address")
	r.Run()
}
