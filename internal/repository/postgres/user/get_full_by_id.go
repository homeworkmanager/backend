package user

import (
	"context"
	"database/sql"
	"errors"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

func (r *Repo) GetFullById(ctx context.Context, id entity.UserID) (entity.UserFullInfo, error) {

	q := "SELECT user_id, name, surname, email, role, group_name FROM users JOIN groups ON users.group_id = groups.group_id " +
		"WHERE user_id = $1"

	var user userFull

	t := r.manager.GetTxOrDefault(ctx)

	err := t.GetContext(ctx, &user, q, id)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.UserFullInfo{}, errs.UserNotFound
	}
	if err != nil {
		return entity.UserFullInfo{}, err
	}
	return user.toUserFull(), nil
}
