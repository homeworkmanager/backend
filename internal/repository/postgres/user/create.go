package user

import (
	"context"
	"github.com/pkg/errors"
	"homeworktodolist/internal/entity"
)

func (r *UserRepo) Create(ctx context.Context, user entity.User) error {
	mName := "UserRepo.Create"

	t := r.manager.GetTxOrDefault(ctx)
	q := "INSERT INTO users(name, surname, email, password, group_id) VALUES ($1, $2, $3, $4, $5);"

	_, err := t.ExecContext(ctx, q, user.Name, user.Surname, user.Email, user.Password, user.GroupId)
	if err != nil {
		return errors.Wrap(err, mName)
	}

	return nil

}
