package main

import (
	"github.com/alaash3lan/paytask/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := app.App{}
	app.Initialize()
	app.Run(":8000")

}
