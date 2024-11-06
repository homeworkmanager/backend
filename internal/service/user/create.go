package user

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"homeworktodolist/internal/entity"
)

type CreateUser struct {
	Name     string
	Surname  *string
	Email    string
	Password string
	GroupId int64
}

func (s *Service) Create(ctx context.Context, req CreateUser) error {
	mName := "UserService.Create"

	user := req.toUser()

	hash, err := s.passwordHash(req.Password)
	if err != nil {
		return errors.Wrap(err, mName)
	}
	user.Password = hash

	if err:= s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, mName)
	}

	return nil

}

func (r *CreateUser) toUser() entity.User {
	return entity.User{
		Name:     r.Name,
		Surname:  r.Surname,
		Email:    r.Email,
		GroupId: r.GroupId,
	}
}

func (s *Service) passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
