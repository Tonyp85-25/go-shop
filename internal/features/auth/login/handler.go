package login

import (
	"example.com/go-shop/internal/config"
	"example.com/go-shop/internal/features/auth/infra"
	"example.com/go-shop/internal/features/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, cfg *config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request
		if err := c.BindJSON(&req); err != nil {
			common.BadRequestResponse(c, "Invalid request data", err)
			return
		}
		userRepo := infra.NewSqlUserRepository(db)

		tokenRepo := infra.NewSqlTokenRepository(db)
		uc := NewUseCase(userRepo, cfg, tokenRepo)

		resp, err := uc.Handle(&req)
		if err != nil {
			common.UnauthorizedResponse(c, "Login failed")
			return
		}
		common.SuccessResponse(c, "User logged in", resp)
	}
}
