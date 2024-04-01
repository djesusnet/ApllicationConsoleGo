package main

import (
	app "ApllicationConsoleGo/Application"
	"log"
	"os"
)

func main() {
	application := app.GenerateApp()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
