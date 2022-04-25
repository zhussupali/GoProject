package main

import (
	"log"
	"twittie"
	"twittie/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(twittie.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("couldn't run server", err.Error())
	}

}