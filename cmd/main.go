package main

import (
	"log"
	"twittie"
	"twittie/pkg/handler"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error init configs: ", err.Error())
	}
	handlers := new(handler.Handler)
	srv := new(twittie.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatal("couldn't run server", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}