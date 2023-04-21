package main

import (
	"finalproj/app"
	"finalproj/config"

	"github.com/joho/godotenv"
)

func init() {

	godotenv.Load()

	err := config.InitGorm()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
