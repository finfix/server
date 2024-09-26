package postgresql

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" //nolint:golint

	"pkg/sql"
)

type PostgreSQLConfig struct {
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1:5433"`
	User     string `env:"SERVER_DB_USER" envDefault:"user"`
	Password string `env:"SERVER_DB_PASSWORD" envDefault:"secret"`
}

func (c *PostgreSQLConfig) GetURL(databaseName string) string {
	return fmt.Sprintf("postgres://%v:%v@%v/%v", c.User, c.Password, c.Host, databaseName)
}

func NewClientSQL(repo PostgreSQLConfig, databaseName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", repo.GetURL(databaseName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db.Unsafe(), nil
}
