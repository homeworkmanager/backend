package user

import "homeworktodolist/internal/tx_manager"

type UserRepo struct {
	manager *tx_manager.TxManager
}

func NewUserRepo(manager *tx_manager.TxManager) *UserRepo {
	return &UserRepo{manager: manager}
}
