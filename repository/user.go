package repository

import (
	"context"

	dbsqlite "github.com/carloscontrerasruiz/gorest/dbSqlite"
	"github.com/carloscontrerasruiz/gorest/models"
)

type UserRepository struct {
	Sqlite *dbsqlite.Sqlite
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (*models.User, error) {
	result, err := r.Sqlite.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) GetById(ctx context.Context, id int64) (*models.User, error) {
	result, err := r.Sqlite.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) Close() error {
	return r.Sqlite.Close()
}

func NewUserRepository() CrudRepository[models.User] {
	return &UserRepository{
		Sqlite: dbsqlite.NewSqliteDB("./basedatos.db"),
	}
}
