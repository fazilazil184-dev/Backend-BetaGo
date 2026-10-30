package main

import (
	"log"

	"github.com/fazilazil184-dev/Backend-BetaGo.git/config"
	"github.com/fazilazil184-dev/Backend-BetaGo.git/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	db := config.Connect()

	r := routes.SetupRouter(db)

	log.Fatal(r.Run(":8080"))

}
