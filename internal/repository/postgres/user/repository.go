package user

import "homeworktodolist/internal/tx_manager"

type Repo struct {
	manager *tx_manager.TxManager
}

func NewUserRepo(manager *tx_manager.TxManager) *Repo {
	return &Repo{manager: manager}
}
