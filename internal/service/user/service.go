package user

import (
	"homeworktodolist/internal/config"
)

type Service struct {
	userRepo      UserRepo
	userRedisRepo UserRedisRepo
	config        *config.Config
}

func NewUserService(userRepo UserRepo, userRedisRepo UserRedisRepo, config *config.Config) *Service {
	return &Service{
		userRepo:      userRepo,
		config:        config,
		userRedisRepo: userRedisRepo,
	}
}
