package database

import (
	"fmt"

	"server/pkg/errors"
	"server/pkg/sql"

	_ "github.com/jackc/pgx/v5/stdlib" //nolint:golint
)

func NewClientSQL(repo RepoConfig, databaseName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%v:%v@%v/%v", repo.User, repo.Password, repo.Host, databaseName))
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db.Unsafe(), nil
}
