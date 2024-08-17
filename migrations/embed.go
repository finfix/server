package migrations

import "embed"

//go:embed pgsql/*.sql
var EmbedMigrationsPostgreSQL embed.FS
