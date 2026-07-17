package infra

import (
	"context"
	"log"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/ecommerce"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestSuite struct {
	suite.Suite
	pgContainer *PostgresContainer
	Db          *gorm.DB
	ctx         context.Context
	Router      *gin.Engine
}

func (s *TestSuite) SetupSuite() {
	s.ctx = context.Background()
	pgContainer, err := NewPostgresContainer(s.ctx)
	if err != nil {
		log.Fatal(err)
	}

	s.pgContainer = pgContainer
	db, err := gorm.Open(postgres.Open(pgContainer.ConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&auth.User{}, &ecommerce.Customer{})
	if err != nil {
		log.Fatal(err)
	}
	s.Db = db
	s.Router = gin.New()

}

func (s *TestSuite) TearDownSuite() {
	if err := s.pgContainer.Terminate(s.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
