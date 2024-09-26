package migrator

import (
	"context"
	"fmt"
	"strings"

	"pkg/log"
)

type migratorLogger struct{}

func newMigratorLogger() migratorLogger {
	return migratorLogger{}
}

func (_ migratorLogger) Fatalf(format string, v ...any) {
	format = strings.TrimSpace(format)
	log.Fatal(context.Background(), fmt.Sprintf(format, v),
		log.ParamsOption("log_source", "goose"),
		log.Skip3PreviousCallersOption(),
	)

}

func (_ migratorLogger) Printf(format string, v ...any) {
	format = strings.TrimSpace(format)
	log.Info(context.Background(), fmt.Sprintf(format, v),
		log.ParamsOption("log_source", "goose"),
		log.Skip3PreviousCallersOption(),
	)
}
