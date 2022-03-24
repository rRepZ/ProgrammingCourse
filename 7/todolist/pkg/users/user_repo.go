package users

import (
	"context"
	"database/sql"
)

type UserData struct {
	Name     string
	Password string
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) Create(ctx context.Context, ud *UserData) (int64, error) {
	res, err := repo.db.ExecContext(ctx, "") // todo SQL-запрос на создание пользователя
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}
