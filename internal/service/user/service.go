package user

import "homeworktodolist/internal/config"

type Service struct {
	userRepo UserRepo
	config      *config.Config
}

func NewUserService(userRepo UserRepo, config *config.Config) *Service {
	return &Service{userRepo: userRepo, config: config}
}
