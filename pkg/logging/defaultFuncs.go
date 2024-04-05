package logging

import (
	"net/http"

	"server/pkg/errors"
)

// DefaultErrorLoggerFunc логгирует ошибки по умолчанию
func DefaultErrorLoggerFunc(err error) {

	customErr := errors.CastError(err)

	logger := GetLogger()

	switch customErr.LogAs {
	case errors.LogAsError:
		logger.Error(err)
	case errors.LogAsWarning:
		logger.Warning(err)
	}
}

func DefaultRequestLoggerFunc(r *http.Request) {
	// Логгируем сообщение по типу METHOD /path
	GetLogger().Info("%v %v", r.Method, r.URL.Path)
}
