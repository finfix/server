package repository

import (
	"context"

	"logger/app/logging"
	pb "logger/app/pblogger"
	"pkg/errors"
	"pkg/sql"
)

type Repository struct {
	dbx    *sql.DB
	logger *logging.Logger
}

func (repo *Repository) AddLog(ctx context.Context, dto *pb.Log) error {

	req := "INSERT INTO logs.logs (path, message, time, log_level, context, service) VALUES (?, ?, ?, ?, ?, ?)"
	if err := repo.dbx.Exec(ctx, req, dto.Path, dto.Message, dto.Time, dto.Level, dto.Context, dto.Service); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}

func New(dbx *sql.DB, logger *logging.Logger) *Repository {
	return &Repository{
		dbx:    dbx,
		logger: logger,
	}
}
