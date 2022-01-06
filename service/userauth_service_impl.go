package service

import (
	"context"
	"database/sql"
	"google_oauth/model"
	"google_oauth/repository"

	"github.com/markbates/goth"
)

type UserAuthServiceImpl struct {
	UserAuthRepo repository.UserAuthRepository
}

func NewUserAuthServiceImpl(db *sql.DB) UserAuthService {
	r := repository.NewUserAuthRepositoryImpl(db)
	return &UserAuthServiceImpl{UserAuthRepo: r}
}

func (s *UserAuthServiceImpl) AddUser(ctx context.Context, user *goth.User) error {
	userModel := model.UserAuth{
		UserId: user.UserID,
		Email:  user.Email,
		Name:   user.Name,
	}

	_, err := s.UserAuthRepo.Add(ctx, &userModel)
	if err != nil {
		return err
	}
	return nil
}
