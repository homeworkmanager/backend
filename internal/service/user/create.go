package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

type CreateUser struct {
	Name     string
	Surname  *string
	Email    string
	Password string
	GroupID  entity.GroupID
}

func (s *Service) Create(ctx context.Context, req CreateUser) error {
	user := req.toUser()

	_, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err == nil {
		return errs.UserExists
	}
	if !(errors.Is(err, errs.UserNotFound)) {
		return err
	}

	hash, err := s.passwordHash(req.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil

}

func (s *Service) passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (r *CreateUser) toUser() entity.User {
	return entity.User{
		Name:    r.Name,
		Surname: r.Surname,
		Email:   r.Email,
		GroupID: r.GroupID,
	}
}
