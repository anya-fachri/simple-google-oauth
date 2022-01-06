package repository

import (
	"context"
	"google_oauth/model"
)

type UserAuthRepository interface {
	Add(ctx context.Context, userauth *model.UserAuth) (*model.UserAuth, error)
}
