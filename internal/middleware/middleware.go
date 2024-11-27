package middleware

type MwManager struct {
	UserRedisRepo UserRedisRepo
}

func NewMwManager(userRedisRepo UserRedisRepo) *MwManager {
	return &MwManager{userRedisRepo}
}
