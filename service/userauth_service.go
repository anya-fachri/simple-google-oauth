package service

import (
	"context"

	"github.com/markbates/goth"
)

type UserAuthService interface {
	AddUser(ctx context.Context, user *goth.User) error
}
