package migrator

import (
	"context"
	"embed"

	"server/app/pkg/errors"
	"server/app/pkg/sql"

	"github.com/pressly/goose/v3"
)

type Migrator interface {
	Up(context.Context) error
	Down(context.Context) error
}

type MigratorConfig struct {
	EmbedMigrations embed.FS // Встроенные файлы миграций
	Dir             string   // Путь к миграциям, так как embedding сохраняет структуру директорий
}

type migrator struct {
	cfg  MigratorConfig
	conn *sql.DB
}

func NewMigrator(conn *sql.DB, config MigratorConfig) Migrator {

	goose.SetBaseFS(config.EmbedMigrations)

	return migrator{
		conn: conn,
		cfg:  config,
	}
}

func (mg migrator) Up(ctx context.Context) error {
	if err := goose.UpContext(ctx, mg.conn.DB.DB, mg.cfg.Dir); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}

func (mg migrator) Down(ctx context.Context) error {
	if err := goose.DownContext(ctx, mg.conn.DB.DB, mg.cfg.Dir); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
