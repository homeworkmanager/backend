package homework

import "homeworktodolist/internal/tx_manager"

type Repo struct {
	manager *tx_manager.TxManager
}

func NewHomeworkRepo(manager *tx_manager.TxManager) *Repo {
	return &Repo{manager: manager}
}
