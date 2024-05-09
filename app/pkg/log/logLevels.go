package log

type logLevel string

// Уровни логов
const (
	warningLevel = logLevel("warn")
	errorLevel   = logLevel("error")
	infoLevel    = logLevel("info")
	fatalLevel   = logLevel("fatal")
	debugLevel   = logLevel("debug")
)
