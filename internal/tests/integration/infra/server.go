package infra

import (
	"example.com/go-shop/internal/config"
	"example.com/go-shop/internal/database"
	applog "example.com/go-shop/internal/logger"
	"example.com/go-shop/internal/server"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTestApp(connectionString string) (router *gin.Engine, db *gorm.DB) {
	log := applog.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	db, err = database.FromString(connectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	gin.SetMode("test")
	srv := server.New(cfg, db, &log)
	router = srv.SetupRoutes()
	return router, db
}
