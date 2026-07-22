package login

import (
	"time"

	"example.com/go-shop/internal/config"
	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/common"
)

type UseCase struct {
	userRepo  auth.UserRepository
	config    *config.JWTConfig
	tokenRepo auth.TokenRepository
}

func NewUseCase(userRepo auth.UserRepository, cfg *config.JWTConfig, tokenRepo auth.TokenRepository) *UseCase {
	return &UseCase{
		userRepo:  userRepo,
		config:    cfg,
		tokenRepo: tokenRepo,
	}
}

func (u *UseCase) Handle(req *Request) (*Response, error) {
	user, err := u.userRepo.FindActive(req.Email)
	if err != nil {
		return nil, common.ErrInvalidCredentials
	}
	if err = auth.CheckPassword(req.Password, user.Password); err != nil {
		return nil, common.ErrInvalidCredentials
	}

	accessToken, refreshToken, err := auth.GenerateTokenPair(u.config, user.PublicID, user.Role)
	if err != nil {
		return nil, err
	}
	tokenModel := auth.RefreshToken{
		UserID:    user.PublicID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(u.config.RefreshTokenExpires),
	}
	if _, err := u.tokenRepo.Create(&tokenModel); err != nil {
		return nil, err
	}

	return &Response{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
