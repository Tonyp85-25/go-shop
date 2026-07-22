package infra

import (
	"context"
	"log"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/ecommerce"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
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
	router, db := CreateTestApp(pgContainer.ConnectionString)

	s.Db = db
	s.Router = router

}

func (s *TestSuite) SetupTest() {
	err := s.Db.AutoMigrate(&auth.User{}, &ecommerce.Customer{}, &auth.RefreshToken{})
	if err != nil {
		log.Fatal(err)
	}
}
func (s *TestSuite) TearDownTest() {
	err := s.Db.Migrator().DropTable(&auth.User{}, &ecommerce.Customer{}, &auth.RefreshToken{})
	if err != nil {
		log.Fatalf("error during test cleanup : %s", err)
	}
}

func (s *TestSuite) TearDownSuite() {
	appDb, err := s.Db.DB()
	if err != nil {
		log.Fatal(err)
	}
	appDb.Close()
	if err := s.pgContainer.Terminate(s.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
