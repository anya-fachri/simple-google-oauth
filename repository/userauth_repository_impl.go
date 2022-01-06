package repository

import (
	"context"
	"database/sql"
	"google_oauth/model"
	"log"
)

type UserAuthRepositoryImpl struct {
	DB *sql.DB
}

func NewUserAuthRepositoryImpl(db *sql.DB) UserAuthRepository {
	return &UserAuthRepositoryImpl{DB: db}
}

func (repo *UserAuthRepositoryImpl) Add(ctx context.Context, userauth *model.UserAuth) (*model.UserAuth, error) {
	queryScript := "SELECT * FROM login_info WHERE email = ?"

	rows, err := repo.DB.QueryContext(ctx, queryScript, userauth.Email)
	if err != nil {
		return userauth, err
	}
	if rows.Next() {
		log.Printf("User %s exists, cancel adding proccess to database", userauth.Email)
		return userauth, nil
	}

	script := "INSERT INTO login_info(userid, email, name) VALUES(?, ?, ?)"

	res, err := repo.DB.ExecContext(ctx, script, userauth.UserId, userauth.Email, userauth.Name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	log.Printf("User %s has been added to database", userauth.Email)
	userauth.Id = int(id)
	return userauth, nil
}
