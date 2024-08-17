package log

import (
	"strconv"
)

type LogLevel uint8

// Уровни логов
const (
	LevelDebug LogLevel = iota + 1
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

func (l LogLevel) ToUpper() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO "
	case LevelWarning:
		return "WARN "
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	}
	return strconv.FormatInt(int64(l), 10)
}

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	}
	return strconv.FormatInt(int64(l), 10)
}
