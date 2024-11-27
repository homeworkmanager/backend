package class

import "homeworktodolist/internal/tx_manager"

type Repo struct {
	manager *tx_manager.TxManager
}

func NewClassRepo(manager *tx_manager.TxManager) *Repo {
	return &Repo{manager: manager}
}
