package database

import (
	"fmt"

	"server/pkg/errors"
	"server/pkg/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewClientSQL(repo RepoConfig, database string) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%v:%v@%v/%v", repo.User, repo.Password, repo.Host, database))
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db.Unsafe(), nil
}
