package main

import (
	"example.com/go-shop/internal/config"
	"example.com/go-shop/internal/database"
	applog "example.com/go-shop/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	log := applog.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	mainDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database connection")
	}
	defer mainDb.Close()

	gin.SetMode(cfg.Server.GinMode)
	log.Info().Msg("Starting server")
}
