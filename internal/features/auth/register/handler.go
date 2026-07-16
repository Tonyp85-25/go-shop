package register

import (
	"example.com/go-shop/internal/features/auth/infra"
	"example.com/go-shop/internal/features/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request
		if err := c.BindJSON(&req); err != nil {
			common.BadRequestResponse(c, "Invalid request data", err)
			return
		}
		userRepo := infra.NewSqlUserRepository(db)
		customerRepo := infra.NewSqlCustomerRepository(db)
		idProvider := infra.UuidProvider{}

		uc := NewUseCase(userRepo, customerRepo, idProvider)
		resp, err := uc.Handle(&req)
		if err != nil {
			common.BadRequestResponse(c, "Registration failed", err)
			return
		}
		common.CreatedResponse(c, "User registered", resp)
	}

}
