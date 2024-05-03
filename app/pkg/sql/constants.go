package sql

import (
	"database/sql"

	"server/app/pkg/errors"
)

var ErrNoRows = sql.ErrNoRows

var secondPathDepthOption = []errors.Option{errors.PathDepthOption(errors.SecondPathDepth)}
