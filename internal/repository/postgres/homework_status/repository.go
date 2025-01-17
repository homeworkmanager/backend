package homework_status

import "homeworktodolist/internal/tx_manager"

type Repo struct {
	manager *tx_manager.TxManager
}

func NewHomeworkStatusRepo(manager *tx_manager.TxManager) *Repo {
	return &Repo{manager: manager}
}
