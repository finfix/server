package sql

import (
	"database/sql"

	"server/pkg/errors"
)

var ErrNoRows = sql.ErrNoRows

var secondPathDepthOption = errors.Options{PathDepth: errors.SecondPathDepth}
