package main

import (
	"log"
	"unisun/api/classroom-listener/src"
	"unisun/api/classroom-listener/src/config"

	"github.com/spf13/viper"
)

func main() {
	envService := config.New("application", "./resources")
	if err := envService.ConfigENV(); err != nil {
		log.Panic(err)
	}
	r := src.App()
	port := viper.GetString("app.port")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
